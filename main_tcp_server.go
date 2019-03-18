package main

<<<<<<< HEAD
import(
	"net"
	"encoding/json"
)

type nameAndEmail struct {
	Name	string
	Email	string
}

func handler(c net.Conn, msg nameAndEmail) {

	response, _ := json.Marshal(msg)

	c.Write([]byte(response))
=======
import (
	"encoding/json"
	"net"
)

type nameAndEmail struct {
	Name string
	Email string
}



func handler(c net.Conn, msg nameAndEmail) {
	response, _ := json.Marshal(msg)
	c.Write([]byte(response))

>>>>>>> 0b5ce10d5d5dc435a2a6b923e40a7f6dbb7123f0
	c.Close()

}

func main() {
	l, err := net.Listen("tcp", ":5000")
<<<<<<< HEAD
	if err != nil {
		panic(err)
	}

	response := nameAndEmail{
		Name:	"Bjørnar",
		Email:	"bjørnar@bjørnar.gmail.com"}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, response)
=======
	if err != nil{
		panic(err) }
	response := nameAndEmail{
		Name: "Bjørnar",
		Email: "bjørnar@bjørnar.gmail.com"}
	for {
		c, err := l.Accept()
		if err!=nil{
			continue }
		go handler(c, response)

>>>>>>> 0b5ce10d5d5dc435a2a6b923e40a7f6dbb7123f0
	}

}