package main

import (

	"fmt"
	"io/ioutil"

	"os"
)

	var f *os.File
func main() {
	f, err := os.Create("proverb.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	f.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	buf,_:=ioutil.ReadFile("proverb.txt")
	fmt.Println(string(buf))

}