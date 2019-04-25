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
<<<<<<< HEAD
	response := jsonify.Encode("Henrik", "henros1996@gmail.com")
=======
	response := jsonify.Encode("trym", "trymerkul@gmail.com")
>>>>>>> 2e5696bc0fb7a952498d3e6e495db195dfbab194
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, response)

	}
}
