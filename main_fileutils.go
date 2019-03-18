package main

//imports fmt for prints
//imports fileutils to access fileutils functions.
import (
	"fmt"

	"./fileutils"
)
<<<<<<< HEAD

var c = "./files/text1.txt"
var d = "./files/text2.txt"
var a = fileutils.FileToByteslice(c)
var b = fileutils.FileToByteslice(d)
=======
var pathA = "./files/text1.txt"
var pathB = "./files/text2.txt"
var a = fileutils.FileToByteslice(pathA)
var b = fileutils.FileToByteslice(pathB)
>>>>>>> eb4f1803950c16f4de5919e9f071a06ea3f1aca4

func main() {
	fmt.Print(a)
	fmt.Println()
	fmt.Print(b)
}
