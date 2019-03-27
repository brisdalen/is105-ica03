package main

import (
	"fmt"
	"net"
	"bufio"
)
// used for 5a and 5b
func main() {
	conn, err := net.Dial("tcp", ":8765")
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\n\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}

