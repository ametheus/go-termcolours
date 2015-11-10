package termcolours

import (
	"fmt"
	"image/color"
)

func Foreground24(c color.Color, s string) string {
	r, g, b, _ := c.RGBA()
	r = (r + 0x80) >> 8
	g = (g + 0x80) >> 8
	b = (b + 0x80) >> 8
	return fmt.Sprintf("%c[38;2;%d;%d;%dm%s%c[0m", ESC, r, g, b, s, ESC)
}

func Background24(c color.Color, s string) string {
	r, g, b, _ := c.RGBA()
	r = (r + 0x80) >> 8
	g = (g + 0x80) >> 8
	b = (b + 0x80) >> 8
	return fmt.Sprintf("%c[48;2;%d;%d;%dm%s%c[0m", ESC, r, g, b, s, ESC)
}
