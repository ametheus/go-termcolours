package main

import (
	"fmt"
	tc "github.com/ametheus/go-termcolours"
)

func main() {
	fmt.Printf("System colours:\n")
	fmt.Printf("%s   %s\n", tc.Black("Black"), tc.Bblack("Bright Black"))
	fmt.Printf("%s    %s\n", tc.Blue("Blue"), tc.Bblue("Bright Blue"))
	fmt.Printf("%s   %s\n", tc.Green("Green"), tc.Bgreen("Bright Green"))
	fmt.Printf("%s    %s\n", tc.Cyan("Cyan"), tc.Bcyan("Bright Cyan"))
	fmt.Printf("%s     %s\n", tc.Red("Red"), tc.Bred("Bright Red"))
	fmt.Printf("%s  %s\n", tc.Purple("Purple"), tc.Bpurple("Bright Purple"))
	fmt.Printf("%s  %s\n", tc.Yellow("Yellow"), tc.Byellow("Bright Yellow"))
	fmt.Printf("%s   %s\n", tc.White("White"), tc.Bwhite("Bright White"))

	fmt.Printf("\n 256ish colour cube\n")
	fmt.Print("4-bit palette: ")
	for i := 0; i < 16; i++ {
		fmt.Print(tc.Foreground8(tc.C256(i), "::"))
	}
	fmt.Print("\n               ")
	for i := 0; i < 16; i++ {
		fmt.Print(tc.Background8(tc.C256(i), "  "))
	}
	fmt.Print("\n")
	for r0 := 0; r0 < 6; r0 += 3 {
		for g := 0; g < 6; g++ {
			for r := 0; r < 3; r++ {
				for b := 0; b < 6; b++ {
					c := tc.Colour256(r+r0, g, b)
					fmt.Print(tc.Foreground8(c, "::"))
				}
				fmt.Print("   ")
			}
			fmt.Print("     ")
			for r := 0; r < 3; r++ {
				for b := 0; b < 6; b++ {
					c := tc.Colour256(r+r0, g, b)
					fmt.Print(tc.Background8(c, "  "))
				}
				fmt.Print("   ")
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
	fmt.Print("4.5-bit greyscale ramp: ")
	for i := 232; i < 256; i++ {
		fmt.Print(tc.Foreground8(tc.C256(i), "::"))
	}
	fmt.Print("\n                        ")
	for i := 232; i < 256; i++ {
		fmt.Print(tc.Background8(tc.C256(i), "  "))
	}
	fmt.Print("\n")
}
