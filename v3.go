package main

import (
	"fmt"
	"time"
)

var (
	s int

)


var mys  = make(map[int]int)
func factoral(n  int){
	//var a int= 2
	for i:=2;i<10000 ;i++{
		//if i%a  != 1 {
		//s = i
		count:=0
		for j :=2;j<i;j++{
			if i%j ==0{
				count++
				break
			}
		}
		if count ==0{
			fmt.Println(i)
	}

	}



}

func main() {
	for i := 0; i < 1;i++{
		go factoral(i)
		time.Sleep(2 * time.Second)
	}

}



