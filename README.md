go-termcolours
==============

ANSI colour escape codes for Go.

Usage
-----

```golang
package main

import (
	"fmt"
	"github.com/ametheus/go-termcolours" tc
)

func main() {
	fmt.Printf( "This is %s; this is %s.\n", tc.Green("green"), tc.Blue("blue") )
}
```
