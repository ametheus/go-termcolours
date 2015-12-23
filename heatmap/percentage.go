package heatmap

type Percentage float64

func (p Percentage) RGBA() (r, g, b, a uint32) {
	a = 0

	if p > 100.0 {
		return 0xffff, 0xffff, 0xffff, 0
	} else if p < 0.0 {
		return 0, 0, 0, 0
	}

	dd := uint32(((p / 100.0) * 0xffff))
	if dd < 0x6000 {
		r = dd * 2
	} else if dd < 0xa000 {
		r = 0xc000 + (dd - 0x6000)
	} else {
		r = 0xffff
	}

	if dd < 0x4000 {
		g = 0
	} else if dd < 0x8000 {
		g = (dd - 0x4000) * 2
	} else if dd < 0xc000 {
		g = 0x4000 + (dd-0x8000)*3
	} else {
		g = 0xffff
	}

	if dd < 0xc000 {
		b = 0
	} else {
		b = (dd - 0xc000) * 4
	}

	return
}
