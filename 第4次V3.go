package main

import (

	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)


func main(){

	for a := 2019210751; a < 2019210790; a++ {
		str := "http://jwzx.node1.cqupt.co/kebiao/kb_stu.php?xh=" + strconv.Itoa(a)
		response, err := http.Get(str)

		if (err != nil) {
			fmt.Println(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		reg := regexp.MustCompile(`201`)
		if reg == nil {
			fmt.Println("MustCompile err")
			return
		}
		buf := string(body)

		reg1 := regexp.MustCompile(`201.......\D\D\D`)
		if reg1 == nil {
			fmt.Println("regexp err")
			return

		}
		result1 := reg1.FindAllStringSubmatch(buf, -1)

		for _,v:=range result1[0]{
			//fmt.Println(v)

			a := string(v)
			reg2 := regexp.MustCompile(`\D\D\D`)
			result2 := reg2.FindAllStringSubmatch(a, -1)
			//fmt.Println(result2)

			for _, l := range result2[0] {
				var c string = string(l)
              fmt.Println(c)

				}
			}
			}

			}


func Tonji (newMap map[string]int)([]string,int) {
	var allcount []int
	var maxcount int
	for _, value := range newMap {
		allcount = append(allcount, value)
	}
	maxcount = allcount[0]
	for i := 0; i < len(allcount); i++ {
		if maxcount < allcount[i] {
			maxcount = allcount[i]
			var maxvalue []string
			for key, value := range newMap {
				if value == maxcount {
					maxvalue = append(maxvalue, key)
				}
			}
			fmt.Println(maxcount)
			return maxvalue, maxcount
		}

		}
var c[]string
	var a int
return c,a
	}
