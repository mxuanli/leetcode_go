package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Students struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (s Human) SayHi() {
	fmt.Printf("名字%s，手机号%s\n", s.name, s.phone)
}

func (e Employee) SayHi() { // 重写SayHi方法
	fmt.Println("重写的方法，公司名：", e.company)
}

func main() {
	s1 := Students{Human{"xuanli", 21, "18303812345"}, "pku"}
	e1 := Employee{Human{"xiaohe", 22, "18303809876"}, "bf"}
	s2 := Students{Human: Human{name: "he", age: 19, phone: "13890877889"}, school: "sg"}

	s1.SayHi()
	s2.SayHi()
	e1.SayHi()
	fmt.Println(e1.phone)
}
