package colour

import "fmt"

const (
	ColourReset = 0
	ColourRed   = iota + 30
	ColourGreen
	ColourYellow
)

func SetColour(c int, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}
