package main

import (
	"fmt"
)

type phone struct {
	brand  string
	cpuNum int
	mem    int
}

type iphone struct {
	phone
	system string
}

func main() {
	p1 := phone{brand: "apple", cpuNum: 10, mem: 2}
	ip1 := iphone{phone{"apple", 4, 2}, "ios"}
	fmt.Println(p1.brand)
	fmt.Println(ip1.cpuNum)

	ip1.setMem()
	ip1.sayBrand()
}

func (p phone) sayBrand() {
	fmt.Println(p.brand, p.mem)
}

func (i *iphone) setMem() {
	i.mem = 10
}
