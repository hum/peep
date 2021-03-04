package peep

import (
  "fmt"
  "net"
  "time"
  "io"

  "github.com/hum/peep/internal"
)

type Whois struct {
  Domains []string
  parser *internal.Parser
  currExt string
}

const (
  IANA_SERVER = "whois.iana.org"
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
    if len(servers) == 0 || servers[0] == "" {
      result, err := w.lookup(name+d, IANA_SERVER)
      if err != nil {
        return false, err
      }
      fmt.Println(result)
    }
  }
  return true, nil
}

func (w *Whois) lookup(name, server string) (string, error) {
  conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, DEFAULT_PORT), time.Second*15)
  if err != nil {
    return "", err
  }
  defer conn.Close()

  conn.SetWriteDeadline(time.Now().Add(time.Second*15))

  payload := []byte(name + "\r\n")
  _, err = conn.Write(payload)
  if err != nil {
    return "", err
  }

  conn.SetReadDeadline(time.Now().Add(time.Second*15))
  response, err := io.ReadAll(conn)
  if err != nil {
    return "", err
  }

  return string(response), nil
}
