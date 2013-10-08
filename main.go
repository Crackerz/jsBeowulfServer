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
	Server.socketServer = &goSocketServer.Server
	Server.pendingJobs = make(chan string,20)
	Server.pendingNodes = make(chan *goSocketServer.Socket,100)
	Server.socketServer.OnConnect(nodeConnected)
	Server.socketServer.OnDisconnect(nodeDisconnected)
	Server.socketServer.OnMessage(nodeMessage)
	go jobWorker(Server.pendingJobs,Server.pendingNodes)
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
