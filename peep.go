package peep

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/hum/peep/internal"
)

type Whois struct {
	Domains []string
	parser  *internal.Parser
	currExt string
}

const (
	IANA_SERVER  = "whois.iana.org"
	DEFAULT_PORT = "43"
)

func initParser() *internal.Parser {
	return &internal.Parser{}
}

func (w *Whois) Search(name string, servers ...string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("Domain name is unspecified.")
	}

	if len(w.Domains) == 0 {
		return false, fmt.Errorf("Domains attr of Whois is not set.")
	}

	if w.parser == nil {
		w.parser = initParser()
	}

	for _, d := range w.Domains {
		// TODO:
		// Allow the input of a specific WHOIS server from param
		if len(servers) == 0 || servers[0] == "" {
			result, err := w.lookup(name+d, IANA_SERVER, time.Second*15)
			if err != nil {
				return false, err
			}
			// TODO:
			// Parse text (internal.Parser) and find if it points to another WHOIS server
			// If not, return; if yes, recursively search for the final one
			fmt.Println(result)
		}
	}
	return true, nil
}

func (w *Whois) lookup(name, server string, timeout time.Second) (string, error) {
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
