package main

import(
	"io/ioutil"
	"os"
	"fmt"
	"github.com/Crackerz/goSocketServer"
	"github.com/Crackerz/fsUtils"
)

type Cluster struct {
	socketServer *goSocketServer.SocketServer
	RootDir string
	program string
	pendingJobs chan string
	pendingNodes chan *Node
}

const (
	program = "program.js"
	pendingDir = "pending"
	processingDir = "processing"
	completedDir = "completed"
)

var Server Cluster

func init() {
	Server.socketServer = &goSocketServer.Server
	Server.socketServer.OnConnect(Server.onConnect)
	Server.socketServer.OnDisconnect(Server.onDisconnect)
	Server.pendingJobs = make(chan string,20)
	Server.pendingNodes = make(chan *Node,100)
	go jobWorker(Server.pendingJobs,Server.pendingNodes)
}

func (c *Cluster) onConnect(s *goSocketServer.Socket) {
	NewNode(s)
	fmt.Printf("Node %d Connected to Server...\n",s.GetId())
}

func (c *Cluster) onDisconnect(s *goSocketServer.Socket) {
	fmt.Printf("Node %d Disconnected from Server...\n",s.GetId())
}

func SetRootDir(folder string) {
	Server.RootDir = folder
	slash := folder+"/"
	initFileSystem(folder)
	SetProgram(slash+program)
	//Begin Monitoring for Changes
	var monitor fsUtils.Monitor
	go monitor.Directory(slash+pendingDir,Server.fileAdd,nil)
}

func (c *Cluster) fileAdd(filename string) {
	fmt.Println("Detected new file: "+filename)
	c.pendingJobs <-filename
}

func initFileSystem(folder string) {
	fmt.Printf("Configuring Directory \"%s/\"...\n",folder)
	dir,err:=ioutil.ReadDir(folder)
	if err!=nil {
		panic("Directory Not Found Error: "+err.Error())
	}

	dirNames:=[]string{pendingDir,processingDir,completedDir}
	dirFound:=[]bool{false,false,false}
	for _,file := range dir {
		for i,name := range dirNames {
			if file.Name()==name&&file.IsDir() {
				fmt.Printf("Found %s...\n",name)
				dirFound[i]=true
			}
		}
	}

	var permissions os.FileMode
	permissions = os.ModeDir | os.ModePerm
	for i,file := range dirNames {
		if !dirFound[i] {
			fmt.Printf("Creating %s...\n",file)
			os.Mkdir(folder+"/"+file,permissions)
		}
	}
}

func SetProgram(filename string) {
	Server.SetProgram(filename)
}

func (c *Cluster) SetProgram(filename string) {
	clientProgram,err:=ioutil.ReadFile(filename)
	if err!=nil {
		panic(err.Error())
	}
	c.program = string(clientProgram)
	fmt.Println(c.program)
}
