package main

import (
	"fmt"
	"time"
)

func main() {

	m1 := make(map[int64]int64)
	m1[1] = 1333331156545664
	m1[2] = 146464989898956
	m1 [3] = 978282061



		for _, v := range m1 {

			fmt.Println("输入的值是", v)
			fmt.Println("输出的值为", time.Unix(100, v))

		}

		}


