package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.43.243:8765")
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\n\n") // Trengs \r for windows?
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}