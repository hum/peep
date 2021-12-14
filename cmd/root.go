package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/hum/peep"
	"github.com/spf13/cobra"
)

var (
	wg *sync.WaitGroup
	in chan string

	domainName string
	domainFile string
	workers int

	whois = peep.Whois{}

	rootCmd = &cobra.Command{
		Use:   "peep",
		Short: "Peep 0.0.1: 🐥 Search for available domains",
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	if len(domainName) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	data, err := getDomains(domainFile)
	if err != nil {
		panic(err)
	}

	wg = new(sync.WaitGroup)
	in = make(chan string, 2*workers)

	domains := strings.Split(string(data), "\n")

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range in {
				if ok, err := whois.Search(domainName, d); err == nil {
					if !ok {
						fmt.Println("Domain: ", domainName+d, " is available.")
					} else {
						fmt.Println("Domain: ", domainName+d, " is taken.")
					}
				}
			}
		}()
	}

	for _, d := range domains {
		in <- d
	}

	close(in)
	wg.Wait()
}

func getDomains(file string) ([]byte, error) {
	if file == "" || len(file) == 0 {
		// File is not present, fetch from github.
		url := "https://raw.githubusercontent.com/hum/peep/main/domains.txt"
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if response.StatusCode == http.StatusOK {
			data, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			return data, nil
		} else {
			return nil, fmt.Errorf("Could not get 'domains.txt' from %s.\nGot response=%s", url, response.Status)
		}
	}

	// Getting local file
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&domainName, "name", "n", "", "domain name to search for")
	rootCmd.Flags().StringVarP(&domainFile, "file", "f", "", "text file containing all of the domains")
	rootCmd.Flags().IntVarP(&workers, "workers", "w", 25, "how many concurrent workers to run")
}
