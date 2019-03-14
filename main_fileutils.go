package main

//imports fmt for prints
//imports fileutils to access fileutils functions.
import (
	"./fileutils"
	"fmt"
)
var c = "./files/text1.txt"
var d = "./files/text2.txt"
var a = fileutils.FileToByteslice(c)
var b = fileutils.FileToByteslice(d)

func main() {
	fmt.Print(a)
	fmt.Println()
	fmt.Print(b)
}
