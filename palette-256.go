package termcolours

import (
	"fmt"
	"image/color"
)

// A 256-colour type. C256 implements image/color.Color
type C256 uint8

func Colour256(r, g, b int) C256 {
	// TODO: actual bounds checking.
	if r < 0 {
		r = 0
	}
	if r > 6 {
		r = 5
	}
	if g < 0 {
		g = 0
	}
	if g > 6 {
		g = 5
	}
	if b < 0 {
		b = 0
	}
	if b > 6 {
		b = 5
	}

	return C256(16 + (r * 36) + (g * 6) + b)
}

func (c C256) RGBA() (r, g, b, a uint32) {
	if c < 16 {
		if c == 0 {
			return 0x2323, 0x3434, 0x3636, 0
		} else if c == 1 {
			return 0xcccc, 0x0000, 0x0000, 0
		} else if c == 2 {
			return 0x4e4e, 0x9a9a, 0x0505, 0
		} else if c == 3 {
			return 0xc4c4, 0xa0a0, 0x0000, 0
		} else if c == 4 {
			return 0x3434, 0x6565, 0xa4a4, 0
		} else if c == 5 {
			return 0x7575, 0x5050, 0x7b7b, 0
		} else if c == 6 {
			return 0x0606, 0x9898, 0x9a9a, 0
		} else if c == 7 {
			return 0xd3d3, 0xd7d7, 0xcfcf, 0
		} else if c == 8 {
			return 0x5555, 0x5757, 0x5353, 0
		} else if c == 9 {
			return 0xefef, 0x2929, 0x2929, 0
		} else if c == 10 {
			return 0x8a8a, 0xe2e2, 0x3434, 0
		} else if c == 11 {
			return 0xfcfc, 0xe9e9, 0x4f4f, 0
		} else if c == 12 {
			return 0x7272, 0x9f9f, 0xcfcf, 0
		} else if c == 13 {
			return 0xadad, 0x7f7f, 0xa8a8, 0
		} else if c == 14 {
			return 0x3434, 0xe2e2, 0xe2e2, 0
		} else if c == 15 {
			return 0xeeee, 0xeeee, 0xecec, 0
		}
	} else if c < 232 {
		b = (uint32((c-16)%6) * 0xffff) / 6
		g = (uint32(((c-16)/6)%6) * 0xffff) / 6
		r = (uint32(((c-16)/36)%6) * 0xffff) / 6
		a = 0
	} else {
		r = (uint32(c-232) * 0xffff) / 24
		g = r
		b = r
		a = 0
	}

	return
}

var all256Colours color.Palette

func init() {
	all256Colours = make([]color.Color, 256)
	for i, _ := range all256Colours {
		all256Colours[i] = C256(i)
	}
}

func adiff(a, b uint32) uint32 {
	if a < b {
		return b - a
	}
	return a - b
}

func Convert256(col color.Color) C256 {
	cc := all256Colours.Convert(col)
	if c2, ok := cc.(C256); ok {
		return c2
	}
	return C256(0)
}

func Foreground8(colour C256, s string) string {
	return fmt.Sprintf("%c[38;5;%dm%s%c[0m", ESC, colour, s, ESC)
}

func Background8(colour C256, s string) string {
	return fmt.Sprintf("%c[48;5;%dm%s%c[0m", ESC, colour, s, ESC)
}
