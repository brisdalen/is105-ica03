package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err1    error
)

func main() {
	newFile, err1 = os.Create("test.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println(newFile)
	newFile.Close()


}
