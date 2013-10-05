package main

import(
	"github.com/Crackerz/goSocketServer"
	"code.google.com/p/go.net/websocket"
)

type Node struct {
	uniq_id int
	socket goSocketServer.Socket
}

func NewNode(socket *goSocketServer.Socket) Node {
	node := Node{socket.GetId(),*socket}
	node.Write("obj = "+Server.program)
	Server.pendingNodes <-&node
	return node
}

func (n *Node) Handle() {
	n.socket.Handle()
}

func (n *Node) Write(text string) {
	n.socket.WriteString(text)
}
