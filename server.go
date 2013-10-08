package main

import(
	"io/ioutil"
	"github.com/Crackerz/goSocketServer"
)

var Server Cluster

type Cluster struct {
	socketServer *goSocketServer.SocketServer
	RootDir string
	program string
	pendingJobs chan string
	pendingNodes chan *goSocketServer.Socket
}

const (
	program = "program.js"
	pendingDir = "pending"
	processingDir = "processing"
	completedDir = "completed"
	resultsDir = "results"
)

func SetProgram(filename string) {
	Server.SetProgram(filename)
}

func (c *Cluster) SetProgram(filename string) {
	clientProgram,err:=ioutil.ReadFile(filename)
	if err!=nil {
		panic(err.Error())
	}
	c.program = string(clientProgram)
}
