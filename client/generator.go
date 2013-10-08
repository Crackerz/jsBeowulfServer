package main

import(	"io/ioutil"
	"strconv")

func main() {
	for i:=0; i < 50; i++ {
		dirname := "pending"
		filename := dirname+"/"+strconv.Itoa(i)+".txt";
		ioutil.WriteFile(filename,[]byte(strconv.Itoa(i)),0777)
	}
}
