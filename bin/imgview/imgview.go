package main

import (
	"flag"
	"fmt"
	tc "github.com/ametheus/go-termcolours"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

var (
	image_file = flag.String("image_file", "", "The image file to display")
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

	var c0, c1 color.Color
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c0 = m.At(x, y)
			c1 = color.Black
			if (y + 1) < bounds.Max.Y {
				c1 = m.At(x, y+1)
			}

			fmt.Print(tc.Background24(c1, tc.Foreground24(c0, BLOCK)))
		}
		fmt.Print("\n")
	}
}
