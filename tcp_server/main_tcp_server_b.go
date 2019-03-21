package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type nameAndEmail struct {
	Name  string
	Email string
}

func handler(c net.Conn, msg nameAndEmail) {
	response, _ := json.Marshal(msg)
	c.Write([]byte(response))
	c.Close()

}
// used in 5b and 5c
func main() {
	fmt.Println("Starting server")
	l, err := net.Listen("tcp", ":8765")
	if err != nil {
		panic(err)
	}
	response := nameAndEmail{
		Name:  "Bjørnar",
		Email: "bjørnar@bjørnar.gmail.com"}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, response)

	}
}