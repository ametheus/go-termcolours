package termcolours

import (
	"fmt"
)

const ESC = '\x1b'

func simple_colour(sequence, input_string string) string {
	return fmt.Sprintf("%c[%sm%s%c[0m", ESC, sequence, input_string, ESC)
}

// 'Normal' colours:
func Black(s string) string {
	return simple_colour("0;30", s)
}
func Blue(s string) string {
	return simple_colour("0;34", s)
}
func Green(s string) string {
	return simple_colour("0;32", s)
}
func Cyan(s string) string {
	return simple_colour("0;36", s)
}
func Red(s string) string {
	return simple_colour("0;31", s)
}
func Purple(s string) string {
	return simple_colour("0;35", s)
}
func Yellow(s string) string {
	return simple_colour("0;33", s)
}
func White(s string) string {
	return simple_colour("0;37", s)
}

// Bright or bold colours:
func Bblack(s string) string {
	return simple_colour("1;30", s)
}
func Bblue(s string) string {
	return simple_colour("1;34", s)
}
func Bgreen(s string) string {
	return simple_colour("1;32", s)
}
func Bcyan(s string) string {
	return simple_colour("1;36", s)
}
func Bred(s string) string {
	return simple_colour("1;31", s)
}
func Bpurple(s string) string {
	return simple_colour("1;35", s)
}
func Byellow(s string) string {
	return simple_colour("1;33", s)
}
func Bwhite(s string) string {
	return simple_colour("1;37", s)
}
