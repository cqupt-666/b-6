package main

import "fmt"

// person 人类
type person struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}

// dove 鸽子
type dove interface {
	gugugu() // 鸽
}

// repeater 复读机
type repeater interface {
	repeat(string) // 复读
}

// repeat 复读
func (p *person) repeat(word string) {
	fmt.Println(word)
}

func (p *person) gugugu() {
	fmt.Println(p.name, "又鸽了")
}

func main() {

	p := &person{
		name:   "xyl",
		age:    18,
		gender: "male",
	}
	p.gugugu()

	var r repeater

	r = p
	r.repeat("helloworld")

var n lingmengjin
n = p
n.shuan("我酸了")
	var a zhenxiangguai
	a = p

a.zhenxiang("真香")


}
type lingmengjin interface {
	shuan(string )
}
func(p*person)shuan(string){
	fmt.Println(p.name ,"我酸了")
}


type zhenxiangguai interface {
	zhenxiang(string)
}
func(p*person)zhenxiang(string ) {
	fmt.Println(p.name, "真香")
}