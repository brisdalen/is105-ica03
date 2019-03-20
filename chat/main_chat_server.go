package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func handler(c net.Conn, msg string) {
	c.Write([]byte(msg))

}

func main() {
	fmt.Println("Starting server")
	l, err := net.Listen("tcp", ":8765")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		// Read input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		go handler(c, text)

	}
}



