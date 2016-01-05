package termcolours

import (
	"os"
)

var TerminalSupports24Bit bool = false

func init() {
	// TODO: sensibly populate bool above
	if os.Getenv("TERM_24BIT") != "" {
		TerminalSupports24Bit = true
	}
}
