package internal

type Parser struct {
  Name string
  Response string
}

func (p *Parser) Cleanup(result string) {
  p.Response = result
  return
}
