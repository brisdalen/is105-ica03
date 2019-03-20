package main

import (
	"fmt"
	"net"
	"io/ioutil"
)

func main() {
<<<<<<< HEAD
	fmt.Println("Client started")
	conn, err := net.Dial("tcp", "192.168.11.108:8765")
=======
	conn, err := net.Dial("tcp", "192.168.43.243:8765")
>>>>>>> 0a1b9b0086c8529cc703f6f4a56913c0a911c067
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	status, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(status))
=======
	fmt.Fprintf(conn, "GET / HTTP/1.0\n\n") // Trengs \r for windows?
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

>>>>>>> 0a1b9b0086c8529cc703f6f4a56913c0a911c067
}