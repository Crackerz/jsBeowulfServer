package main

import (
	"fmt"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

//Establish all http listeners
func init() {
	fmt.Printf("Configuring Server...\n")
	http.HandleFunc("/",website)
	http.Handle("/socket",websocket.Handler(socket))
	SetRootDir("client")
}

func main() {
	fmt.Printf("Starting Server...\n")
	http.ListenAndServe(":8080",nil)
}

func website(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "website")
}

func socket(ws *websocket.Conn) {
	fmt.Printf("Received Socket Connection...\n")
	n:=NewNode(ws)
	n.Socket.Handle()
	fmt.Printf("Handled Connection\n")
}
