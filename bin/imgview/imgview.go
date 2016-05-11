package main

import (
	"flag"
	"fmt"
	tc "github.com/thijzert/go-termcolours"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

var (
	image_file = flag.String("image_file", "", "The image file to display")
	use_24bit  = flag.Bool("use_24bit", true, "Use 24-bit colours")
)

const BLOCK = "\xe2\x96\x80"

func init() {
	flag.Parse()
}

func main() {
	reader, err := os.Open(*image_file)
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

	if *use_24bit {
		Write24(m, bounds)
	} else {
		Write8(m, bounds)
	}
}

func Write24(i image.Image, bounds image.Rectangle) {
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

func Write8(i image.Image, bounds image.Rectangle) {
	var c0, c1 tc.C256
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c0 = iget(i, bounds, x, y)
			c1 = 232
			if (y + 1) < bounds.Max.Y {
				c1 = iget(i, bounds, x, y+1)
			}

			fmt.Print(tc.Background8(c1, tc.Foreground24(c0, BLOCK)))
		}
		fmt.Print("\n")
	}
}
