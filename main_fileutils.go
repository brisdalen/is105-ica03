package main

//imports fmt for prints
//imports fileutils to access fileutils functions.
import (
	"fmt"

	"github.com/Henrikohlsen/Is-105/is105-ica03/fileutils"
)

var pathA = "./files/text1.txt"
var pathB = "./files/text2.txt"
var a = fileutils.FileToByteslice(pathA)
var b = fileutils.FileToByteslice(pathB)

func main() {
	fmt.Print(a)
	fmt.Println()
	fmt.Print(b)
}
