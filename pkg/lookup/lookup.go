package lookup

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func IsTaken(dm string) (taken bool, err error) {
	// Query actual WHOIS information
	whois, err := whois.Whois(dm)
	if err != nil {
		return
	}

	// Parse the data into a struct
	_, err = whoisparser.Parse(whois)
	if err != nil {
		// Naively only check if the domain is not found
		// for now we ignore any other case
		if err == whoisparser.ErrNotFoundDomain {
			return false, nil
		} else if err == whoisparser.ErrReservedDomain {
			return false, nil
		}
		return
	}
	// Is taken because we got WHOIS data back
	return true, nil
}
