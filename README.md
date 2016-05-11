go-termcolours
==============

ANSI colour escape codes for Go.

Installation
------------

```
go get -u github.com/thijzert/go-termcolours
```

Usage
-----

```golang
package main

import (
	"fmt"
	"github.com/thijzert/go-termcolours" tc
)

func main() {
	fmt.Printf( "This is %s; this is %s.\n", tc.Green("green"), tc.Blue("blue") )
}
```

License
-------

This repository is shared under a BSD 3-clause license. See the file COPYING for details.
