package main

import(	"fmt"
	"io/ioutil"
	"strings")

func jobWorker(jobs chan string, nodes chan *Node) {
	//A worker's job is never done
	for {
		filename := <-jobs
		node := <-nodes
		fmt.Printf("Sending %s to Node %d\n",filename,node.uniq_id)
		rawdata,err:=ioutil.ReadFile(Server.RootDir+"/"+pendingDir+"/"+filename)
		strdata := strings.TrimSpace(string(rawdata))
		if err!=nil {
			panic(err.Error())
		}
		node.Write(strdata)
	}
}
