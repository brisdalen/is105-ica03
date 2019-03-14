package fileinfo

import (
	"fmt"
	"log"
	"os"
)



var (
	fileinfo os.FileInfo
	err      error
)
var filename string
var read = "./is105-ica03/files/text1.txt"
//this is a function to read files
func FileReader() {
	fileinfo, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("file name:", fileinfo.Name())
	fmt.Println("Size in bytes", fileinfo.Size())
	fmt.Println("Permissions", fileinfo.Mode())
	fmt.Println("Last modified", fileinfo.ModTime())
	fmt.Println("Is a directory", fileinfo.IsDir())
	fmt.Printf("System Interface type: %T\n", fileinfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileinfo.Sys())

}
