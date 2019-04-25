package main

import (
	"fmt"
	"net"
	"../jsonify"
)

func handler(c net.Conn, msg []byte) {
	c.Write([]byte(msg))
	c.Close()

}
// used in 5b and 5c
func main() {
	fmt.Println("Starting server")
	l, err := net.Listen("tcp", ":8765")
	if err != nil {
		panic(err)
	}
	response := jsonify.Encode("Henrik", "henros1996@gmail.com")
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, response)

	}
}
