package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileinfo os.FileInfo
	err	error
)
var testRead = "./files/text1.txt"
var testRead1 = "./files/text2.txt"
	func main () {

		fileinfo, err = os.Stat("./files/pg100.txt")
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("File name:", fileinfo.Name())
		fmt.Println("Size in bytes:", fileinfo.Size())
		fmt.Println("Permissions:", fileinfo.Mode())
		fmt.Println("Last modified", fileinfo.ModTime())
		fmt.Println("Is Directory:", fileinfo.IsDir())
		fmt.Printf("System interface type: %T\n", fileinfo.Sys())
		fmt.Printf("System info: %+v\n\n", fileinfo.Sys())

	}
