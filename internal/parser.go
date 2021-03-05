package internal

import (
	"fmt"
	"strings"
)

const (
	REFER = "refer:"
	WHOIS = "whois:"
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
	split := strings.Split(strings.TrimSpace(data), "\n")
	// special case for domains like .ai
	if strings.Contains(split[0], "Domain Name: ") {
		data = split[1]
	} else {
		data = split[0]
	}

	for _, v := range keywords {
		if strings.Contains(strings.ToLower(data), v) {
			return false
		}
	}
	fmt.Println(p.Domain)
	return true
}

func (p *Parser) getReferServer(data string) (string, error) {
	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, REFER) || strings.Contains(line, WHOIS) {
			return line, nil
		}
	}
	return "", fmt.Errorf("Could not find ref server.")
}
