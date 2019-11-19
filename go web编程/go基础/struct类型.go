package main

import (
	"fmt"
)

//定义一个 "人" 类型，有名字和年龄两个属性
type person struct {
	name string
	age  int
}

func main() {
	var p person
	p.name = "xuanli"
	p.age = 21
	fmt.Println("person name is", p.name)
	var p1 = person{"xiaohe", 21}
	fmt.Println("person name is", p1.name)
}
