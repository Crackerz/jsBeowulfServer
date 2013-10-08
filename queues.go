package main

import(	"fmt"
	"io/ioutil"
	"github.com/Crackerz/goSocketServer")

var nodeToJob map[int]string

func jobWorker(jobs chan string, nodes chan *goSocketServer.Socket) {
	nodeToJob = make(map[int]string)
	for { //A worker's job is never done
		fmt.Println("Searching for file...")
		filename := <-jobs
		fmt.Println("Found file: ",filename)
		data,err:=ioutil.ReadFile(Server.RootDir+"/"+pendingDir+"/"+filename)
		if err!=nil {
			fmt.Println(err.Error()+" Trying new file...")
		} else {
			fmt.Println("Pairing file with node...")
			node := <-nodes
			markProcessing(filename)
			nodeToJob[node.GetId()] = filename
			s:=goSocketServer.Socket(*node)
			fmt.Printf("Sending %s to Node %d\n",filename,s.GetId())
			node.SendBytes(data)
		}
	}
}
