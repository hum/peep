package colour

import (
	"testing"
)

func TestColourSetForString(t *testing.T) {
	tests := []struct {
		colour         int
		input          string
		expectedOutput string
	}{
		{
			ColourRed, "hello, world!", "\x1b[31mhello, world!\x1b[0m",
		},
		{
			ColourGreen, "test", "\x1b[32mtest\x1b[0m",
		},
		{
			ColourYellow, "test", "\x1b[33mtest\x1b[0m",
		},
	}

	for _, tt := range tests {
		result := SetColour(tt.colour, tt.input)
		if result != tt.expectedOutput {
			t.Fatalf("output: %s does not match expected output: %s", result, tt.expectedOutput)
		}
	}
}
