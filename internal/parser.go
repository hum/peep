package internal

type Parser struct {
	Name     string
	Response string
}

// TODO

func (p *Parser) Cleanup(result string) {
	// TODO:
	// Parse text recieved from the WHOIS server
	p.Response = result
	return
}
