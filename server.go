package main

import(
	"io/ioutil"
	"os"
	"fmt"
	"github.com/crackerz/goSocketServer"
)

type Cluster struct {
	socketServer goSocketServer.SocketServer
	RootDir string
	program string
}

const (
	program = "program.js"
	pendingDir = "pending"
	processingDir = "processing"
	completedDir = "completed"
)

var Server Cluster

func init() {
	Server.socketServer = goSocketServer.Server
}

func SetRootDir(folder string) {
	slash := folder+"/"
	initFileSystem(folder)
	SetProgram(slash+program)
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
	(*c).program = string(clientProgram)
}
