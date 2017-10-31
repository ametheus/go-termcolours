package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"syscall"

	"github.com/nfnt/resize"
	"golang.org/x/crypto/ssh/terminal"

	tc "github.com/thijzert/go-termcolours"
)

var (
	text_aspect  = flag.Float64("text_aspect", 0.944444, "Aspect ratio for your terminal font")
	use_24bit    = flag.Bool("use_24bit", false, "Use 24-bit colours")
	force_width  = flag.Int("width", 0, "Force output width")
	force_height = flag.Int("height", 0, "Force output height")
)

const BLOCK = "\xe2\x96\x80"

func init() {
	flag.Parse()
}

func main() {
	var err error
	termx, termy := *force_width, *force_height
	if termx == 0 && termy == 0 {
		termx, termy, err = terminal.GetSize(syscall.Stdout)
		if err != nil {
			log.Print(err)
			termx, termy = 80, 25
		}
	} else {
		if termx == 0 {
			termx = termy * 1000
		} else if termy == 0 {
			termy = termx * 1000
		}
	}

	// We can stack two pixels in one character
	termy *= 2

	for _, image_file := range flag.Args() {
		reader, err := os.Open(image_file)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()

		m, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		bounds := m.Bounds()
		fmt.Printf("Image is %s by %s pixels wide\n", tc.Green(fmt.Sprintf("%d", bounds.Max.X)), tc.Green(fmt.Sprintf("%d", bounds.Max.Y)))

		nx, ny := boundbox(bounds.Max.X, bounds.Max.Y, termx, termy)
		mm := resize.Resize(uint(nx), uint(ny), convertRGBA(m), resize.Lanczos3)

		bounds = mm.Bounds()
		fmt.Printf("Image is now %s by %s pixels wide\n", tc.Green(fmt.Sprintf("%d", bounds.Max.X)), tc.Green(fmt.Sprintf("%d", bounds.Max.Y)))

		if *use_24bit {
			Write24(mm)
		} else {
			Write8(mm)
		}
	}
}

func convertRGBA(in image.Image) image.Image {
	if m, ok := in.(*image.RGBA); ok {
		return m
	}

	bounds := in.Bounds()

	m := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, in.At(x, y))
		}
	}
	return m
}

func boundbox(imgx, imgy, bx, by int) (x, y int) {
	if imgx < 1 || imgy < 1 || bx < 1 || by < 1 {
		return 1, 1
	}

	term_aspect := float64(by) / float64(bx)
	aspect := *text_aspect * (float64(imgy) / float64(imgx))

	if aspect >= term_aspect {
		y = by
		x = int((float64(by) / aspect) + 0.5)
		if x > bx {
			x = bx
		}
	} else {
		x = bx
		// We can stack two pixels in one character
		y = 2 * int((float64(bx)*aspect*0.5)+0.5)
		if y > by {
			y = by
		}
	}
	return
}

func Write24(i image.Image) {
	bounds := i.Bounds()
	var c0, c1 color.Color
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c0 = i.At(x, y)
			c1 = color.Black
			if (y + 1) < bounds.Max.Y {
				c1 = i.At(x, y+1)
			}

			fmt.Print(tc.Background24(c1, tc.Foreground24(c0, BLOCK)))
		}
		fmt.Print("\n")
	}
}

func cdiff(before color.Color, after tc.C256) (r, g, b int32) {
	r0, g0, b0, _ := before.RGBA()
	r1, g1, b1, _ := after.RGBA()

	r = int32(r1) - int32(r0)
	g = int32(g1) - int32(g0)
	b = int32(b1) - int32(b0)
	return
}

func pos(a, b int32) uint32 {
	a += b
	if a < 0 {
		return 0
	}
	return uint32(a)
}

func iadd(i image.Image, x, y int, dR, dG, dB int32, multiplier float64) {
	col := i.At(x, y)

	r, g, b, _ := col.RGBA()
	r = pos(int32(r), int32(float64(dR)*multiplier))
	g = pos(int32(g), int32(float64(dG)*multiplier))
	b = pos(int32(b), int32(float64(dB)*multiplier))

	cnew := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 0}
	i.(*image.RGBA).Set(x, y, cnew)
}

func iget(i image.Image, bounds image.Rectangle, x, y int) tc.C256 {
	col := i.At(x, y)
	aft := tc.Convert256(col)

	dr, dg, db := cdiff(col, aft)

	if (x + 1) < bounds.Max.X {
		iadd(i, x+1, y, dr, db, dg, 7.0/16.0)
		if (y + 1) < bounds.Max.Y {
			iadd(i, x+1, y+1, dr, db, dg, 1.0/16.0)
		}
	}
	if (y + 1) < bounds.Max.Y {
		iadd(i, x, y+1, dr, db, dg, 5.0/16.0)
		if (x - 1) >= bounds.Min.X {
			iadd(i, x-1, y+1, dr, db, dg, 3.0/16.0)
		}
	}

	return aft
}

func Write8(i image.Image) {
	var c0, c1 tc.C256
	bounds := i.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c0 = iget(i, bounds, x, y)
			c1 = 232
			if (y + 1) < bounds.Max.Y {
				c1 = iget(i, bounds, x, y+1)
			}

			fmt.Print(tc.Background8(c1, tc.Foreground8(c0, BLOCK)))
		}
		fmt.Print("\n")
	}
}
