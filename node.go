package main

import(
	"github.com/Crackerz/goSocketServer"
	"fmt"
)

func nodeConnected(s *goSocketServer.Socket) {
	fmt.Printf("Node %d Connected to Server...\n",s.GetId())
	s.SendBytes([]byte(Server.program))
	Server.pendingNodes<-s
}

func nodeDisconnected(s *goSocketServer.Socket) {
	fmt.Printf("Node %d Disconnected from Server...\n",s.GetId())
}

func nodeMessage(s *goSocketServer.Socket, message []byte) {
	fmt.Println("Received Message: ", string(message)," From: ",s.GetId())
	markComplete(nodeToJob[s.GetId()],message)
	delete(nodeToJob,s.GetId())
	Server.pendingNodes<-s
}
