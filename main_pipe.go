package main

import (
	"fmt"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe"
	"io/ioutil"
	"os"
)

func main() {
	text, err := ioutil.ReadFile("files/pg100.txt")

	ts, _ := os.Stat("files/pg100.txt")
	size := ts.Size()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if size <= 100000 {
		PrintValues(text)
	} else {
		fmt.Println("The file is currently to large")
	}
}
func PrintValues(b []byte) {
	hex := pipe.HexReturn(string(b))
	fmt.Println(hex, "\n", "This Hexcode length is ", len(hex), "in bytes", "\n")
	b64 := pipe.Base64Return(hex)
	fmt.Println(b64, "\n", "This Base64 length is", len(b64), "in bytes", "\n")
	gzip := pipe.GZipReturn(b64)
	fmt.Println(gzip, "\n", "This GZIP length is", len(gzip), "in bytes")
}
