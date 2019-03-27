
package main

import (
	"fmt"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe"
)

func main() {
	hex := pipe.HexReturn("Hallo")
	fmt.Println(hex, "\n")
	b64 := pipe.Base64Return(hex)
	fmt.Println(b64, "\n")
	gzip := pipe.GZipReturn(b64)
	fmt.Println(gzip)
	}


