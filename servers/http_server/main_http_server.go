package main

import(
	"fmt"
	"net/http"
	"../jsonify"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := jsonify.Encode("bjørnar", "bjørnar@bjørnar.gmail.com")
	fmt.Fprintf(w, string(msg))
}

func main() {
	http.HandleFunc("/connected", handler)
	http.ListenAndServe(":8765", nil)
}

