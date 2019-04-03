package main

import (
	"github.com/Nosp1/Is-105/is105-ica03/fileinfo"
	"os"
)

var testRead = "./files/text1.txt"
var testRead1 = "./files/text2.txt"

func main() {
	if len(os.Args) == 2 {
		//FileReader(os.Args[1])
		feilmelding := "Du mangler -f"
		panic(feilmelding)
	}
	if len(os.Args) == 3 {
		param := os.Args[1]

		if param == "-f" {
			fileinfo.FileReader(os.Args[2])
		}
	}

}
