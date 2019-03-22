package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)
// used for 5c
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Did you forget the IP address?")
	} else {
		ip := os.Args[1] + ":8765"
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			fmt.Println("Error")
		}

		fmt.Fprintf(conn, "GET / HTTP/1.0\n\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status)
	}

}

