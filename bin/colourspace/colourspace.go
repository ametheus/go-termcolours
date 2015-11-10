package main

import (
	"fmt"
	tc "github.com/ametheus/go-termcolours"
)

func main() {
	fmt.Printf("%s  %s\n", tc.Black("Black"), tc.Bblack("Bright Black"))
	fmt.Printf("%s  %s\n", tc.Blue("Blue"), tc.Bblue("Bright Blue"))
	fmt.Printf("%s  %s\n", tc.Green("Green"), tc.Bgreen("Bright Green"))
	fmt.Printf("%s  %s\n", tc.Cyan("Cyan"), tc.Bcyan("Bright Cyan"))
	fmt.Printf("%s  %s\n", tc.Red("Red"), tc.Bred("Bright Red"))
	fmt.Printf("%s  %s\n", tc.Purple("Purple"), tc.Bpurple("Bright Purple"))
	fmt.Printf("%s  %s\n", tc.Yellow("Yellow"), tc.Byellow("Bright Yellow"))
	fmt.Printf("%s  %s\n", tc.White("White"), tc.Bwhite("Bright White"))
}
