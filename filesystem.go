package main

import(	"fmt"
	"io/ioutil"
	"os"
	"github.com/Crackerz/fsUtils"
)

func initFileSystem(folder string) {
	fmt.Printf("Configuring Directory \"%s/\"...\n",folder)
	dir,err:=ioutil.ReadDir(folder)
	if err!=nil {
		panic("Directory Not Found Error: "+err.Error())
	}

	dirNames:=[]string{pendingDir,processingDir,completedDir,resultsDir}
	dirFound:=[]bool{false,false,false,false}
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

func SetRootDir(folder string) {
	Server.RootDir = folder
	slash := folder+"/"
	initFileSystem(folder)
	SetProgram(slash+program)
	//Begin Monitoring for Changes
	var monitor fsUtils.Monitor
	go monitor.Directory(slash+pendingDir,fileAdd,nil)
}

func fileAdd(filename string) {
	fmt.Println("Detected new file: "+filename)
	Server.pendingJobs <-filename
}

func markProcessing(filename string) {
	src:=Server.RootDir+"/"+pendingDir+"/"+filename
	dst:=Server.RootDir+"/"+processingDir+"/"+filename
	fmt.Println("Moving from: ",src," to ",dst)
	err:=os.Rename(src,dst)
	if err!=nil {
		fmt.Println(err.Error())
	}
}

func markComplete(filename string, result []byte) {
	src:=Server.RootDir+"/"+processingDir+"/"+filename
	dst:=Server.RootDir+"/"+completedDir+"/"+filename
	fmt.Println("Moving from: ",src," to ",dst)
	err:=os.Rename(src,dst)
	if err!=nil {
		fmt.Println(err.Error())
	}

	out:=Server.RootDir+"/"+resultsDir+"/"+filename
	ioutil.WriteFile(out,result,os.ModePerm)
}
