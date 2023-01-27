package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/hum/peep/pkg/colour"
	"github.com/hum/peep/pkg/domain"
	"github.com/hum/peep/pkg/lookup"
)

var (
	domainName     string
	workerPool     int
	runningWorkers int
	wg             sync.WaitGroup

	availableCount int
	takenCount     int
)

const (
	IS_TAKEN     = "[❌] "
	IS_AVAILABLE = "[✅] "
)

func resolve(in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		runningWorkers--
	}()

	for dn := range in {
		taken, err := lookup.IsTaken(dn)
		if err != nil {
			fmt.Printf("error for domain: %s, got error: %s\n", dn, err)
			continue
		}

		var result string
		if taken {
			// Domain is taken
			result = fmt.Sprintf("%s %s", IS_TAKEN, dn)
			result = colour.SetColour(colour.ColourRed, result)
			takenCount++
		} else {
			// Domain is free
			result = fmt.Sprintf("%s %s", IS_AVAILABLE, dn)
			result = colour.SetColour(colour.ColourGreen, result)
			availableCount++
		}
		out <- result
	}
}

func main() {
	var maxWorkerPool int = runtime.GOMAXPROCS(0) * 2

	flag.StringVar(&domainName, "domain", "", "domain name to look up")
	flag.StringVar(&domainName, "d", "", "domain name to look up")
	flag.IntVar(&workerPool, "jobs", maxWorkerPool, "sets the amount of coroutines")
	flag.IntVar(&workerPool, "j", maxWorkerPool, "sets the amount of coroutines")
	flag.Parse()

	if domainName == "" {
		fmt.Println("no domain name provided")
		flag.Usage()
		return
	}

	in := make(chan string, len(domain.Extensions))
	out := make(chan string, len(domain.Extensions))

	// Initialise worker pool
	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		runningWorkers++

		go resolve(in, out, &wg)
	}

	fmt.Printf("Searching TLDs for domain name: %s\n", colour.SetColour(colour.ColourYellow, domainName))
	start := time.Now()

	for _, dn := range domain.Extensions {
		in <- fmt.Sprintf("%s%s", domainName, dn)
	}
	// Close input channel since all inputs have been sent already
	close(in)

	// Inline goroutine to print out the output
	wg.Add(1)
	go func(output <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			select {
			case result := <-output:
				fmt.Println(result)
			default:
				// Once all workers are finished, there won't be any more results
				if runningWorkers == 0 {
					return
				}
			}
		}
	}(out, &wg)

	wg.Wait()
	close(out)

	end := time.Now()
	dt := end.Sub(start)

	response := fmt.Sprintf("query took: %s, performed %d queries -- %d available, %d taken",
		dt,
		len(domain.Extensions),
		availableCount,
		takenCount,
	)

	fmt.Println(colour.SetColour(colour.ColourYellow, response))
}
