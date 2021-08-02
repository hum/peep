package internal

import (
	"fmt"
	"strings"
)

const (
	WHOIS = "whois:"
  REGISTRAR = "Registrar WHOIS Server:"
)

var (
	keywords = []string{
		"no match for",
		"no data found",
		"not found",
		"does not exist",
		"no matching",
		"is free",
		"is available for registration",
		"has not been registered",
		"no object found",
		"available",
		"no entries found",
		"No_Se_Encontro_El_Objeto/Object_Not_Found",
		"Status: invalid",
	}
)

type Parser struct {
	Domain string
}

func (p *Parser) GetReferServer(data string) (string, error) {
	refer, err := p.getReferServer(data)
	if err != nil {
		return "", err
	}

	// data currently should be:
	// "refer:    whois.server.domain"
	return strings.TrimSpace(refer[6:]), nil
}

func (p *Parser) IsFound(data string) bool {
	var split []string

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		if len(line) <= 1 {
			continue
		}
		/*if strings.Contains(string(line[0]), "%") || strings.Contains(string(line[0]), "#") {
		  if strings.Contains(string(line[1]), "%") {
		    split = append(split, line)
		  } else {
		    continue
		  }
		}*/
		split = append(split, line)
	}
	// special case for domains like .ai
	// TODO:
	// rewrite this horrid thing
	// but hey, it works for now
	for _, line := range split {
		for _, k := range keywords {
			if strings.Contains(strings.ToLower(line), strings.ToLower(k)) {
				return false
			}
		}
	}
	//fmt.Println(p.Domain)
	return true

	/*if strings.Contains(split[0], "Domain Name: ") || strings.Contains(split[0], "Domain: ") {
	    if strings.Contains(split[1], "Script: ") {
	      data = split[2]
	    } else {
	      data = split[1]
	    }
		} else {
			data = split[0]
		}

		for _, v := range keywords {
			if strings.Contains(strings.ToLower(data), v) {
				return false
			}
		}
		fmt.Println(p.Domain)
		return true*/
}

func (p *Parser) getReferServer(data string) (string, error) {
	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, WHOIS) || strings.Contains(line, REGISTRAR) {
      return line, nil
		}
	}
  return "", fmt.Errorf("Could not find ref server.")
}
