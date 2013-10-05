package main

import (
	"fmt"
	"net/http"
	"code.google.com/p/go.net/websocket"
	"github.com/Crackerz/goSocketServer"
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
	s:=goSocketServer.NewSocket(ws)
	s.Handle()
}
