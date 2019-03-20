package main

import (
	"fmt"
	"net"
	"io/ioutil"
)

func main() {
	fmt.Println("Client started")
	conn, err := net.Dial("tcp", "192.168.11.108:8765")
	if err != nil {
		panic(err)
	}

	status, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(status))
}