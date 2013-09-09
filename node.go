package main

import(
	"github.com/crackerz/goSocketServer"
	"code.google.com/p/go.net/websocket"
)

type Node struct {
	uniq_id int
	Socket goSocketServer.Socket
}

func NewNode(ws *websocket.Conn) Node {
	socket:=goSocketServer.NewSocket(ws)
	return Node{socket.GetId(),socket}
}
