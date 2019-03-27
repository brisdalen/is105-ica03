package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//Open file and create a buffered reader on top
	file, err := os.Open("pg100.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)

	byteSlice := make([]byte, 100)
	byteSlice, err = bufferedReader.Peek(100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("peeked at 100 bytes: %s\n", byteSlice)

	// Read and advance pointer
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// Ready 1 byte. Error if no byte to read
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// Read up to and including delimiter
	// Returns byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

}
