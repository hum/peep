package lookup

import "testing"

func TestLookupTakenDomain(t *testing.T) {
	tests := []struct {
		domainName string
	}{
		{"google.io"},
		{"youtube.com"},
	}

	for _, tt := range tests {
		taken, err := IsTaken(tt.domainName)
		if err != nil {
			t.Fatalf("unexpected error in IsTaken test: %s", err)
		}

		if !taken {
			t.Fatalf("domain is reported as available while taken: %s", tt.domainName)
		}
	}
}

func TestLookupAvailableDomain(t *testing.T) {
	tests := []struct {
		domainName string
	}{
		{"akljsdankj123sadnj.io"},
		{"xczzxjnkals1sx.com"},
	}

	for _, tt := range tests {
		taken, err := IsTaken(tt.domainName)
		if err != nil {
			t.Fatalf("unexpected error in IsTaken test: %s", err)
		}

		if taken {
			t.Fatalf("domain is reported as taken while should be available: %s", tt.domainName)
		}
	}

}
