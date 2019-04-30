package main

import (
	"fmt"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe/b64"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe/gzip"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe/hex"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("No extra arguments provided.")
	}

	file, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	x := hex.HexReturnBytes(file)
	b := b64.Base64ReturnBytes(file)
	gz := gzip.GZipReturnBytes(b)
	result := hex.HexReturnBytes(gz)

	i1 := len(x)
	i2 := len(b)
	i3 := len(gz)

	LimitPrintln(x, 10, " - hexx")
	LimitPrintln(b, 10, "- base64")
	LimitPrintln(gz, 10, "- gzip")
	LimitPrintln(result, 10, "- base64 + gzip in hex")

	fmt.Println("-------------------------------------------------------------")

	fmt.Println(len(x), ": Original length in hex-encoding")
	fmt.Println(len(b), ": Length of the input file in base64-encoding")
	fmt.Println(len(gz), ": Length of the input file after base64- and gzip-encoding")
	fmt.Println(len(result), ": Length of input after pipe in hex-encoding")

	fmt.Println("-------------------------------------------------------------")

	fmt.Println("The Base64 Compression rate is : ", CompressionRatio(i1, i2), " %")
	fmt.Println("The Gzip Compression rate is : ",  CompressionRatio(i1, i3), " %")

}

func LimitPrintln(b []byte, limit int, s string) {
	fmt.Println(b[:limit], ": The first", limit, "bytes.", s)
}



func CompressionRatio(old, new int) float64  {
	oldn:= float64(old)
	newn := float64(new)
	percentage := 100 * (1 - (newn / oldn))
	return percentage
	}