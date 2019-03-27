package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.Reader
	s := ioutil.ReadAll("test.txt")
	fmt.Printf("%X", s)
}
