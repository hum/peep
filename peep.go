package peep

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/hum/peep/internal"
)

type Whois struct {
	parser *internal.Parser
}

const (
	IANA_SERVER  = "whois.iana.org"
	DEFAULT_PORT = "43"
)

func initParser() *internal.Parser {
	return &internal.Parser{}
}

func (w *Whois) Search(name, domain string, servers ...string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("Domain name is unspecified.")
	}
	if w.parser == nil {
		w.parser = initParser()
	}

	/*
	   TODO:
	   Allow the input of a specific WHOIS server from param
	*/
	if len(servers) == 0 || servers[0] == "" {
		result, err := w.lookup(name+domain, IANA_SERVER, time.Second*15)
		if err != nil {
			return false, err
		}
		result = cleanup(result)

		w.parser.Domain = domain
		/*
		   TODO:
		   Parse response (internal.Parser) and find if it points to another WHOIS server
		   If not, return; if yes, search for the final one
		*/
		ref, err := w.parser.GetReferServer(result)
		if err != nil {
			return false, err
		}

		result, err = w.lookup(name+domain, ref, time.Second*15)
		if err != nil {
			return false, err
		}

		if ok := w.parser.IsFound(result); !ok {
			return false, nil
		}
	}
	return true, nil
}

// temp function to de-clutter output
func cleanup(data string) string {
	var result []string

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 || string(line[0]) == "%" {
			continue
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

func (w *Whois) lookup(name, server string, timeout time.Duration) (string, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, DEFAULT_PORT), timeout)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(timeout))
	padding := "\r\n"

	payload := []byte(name + padding)
	_, err = conn.Write(payload)
	if err != nil {
		return "", err
	}

	conn.SetReadDeadline(time.Now().Add(timeout))
	response, err := io.ReadAll(conn)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
