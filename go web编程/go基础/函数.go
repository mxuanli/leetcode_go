package main

import (
	"fmt"
)

type intP *int

func main() { // 函数穿参穿的是参数的一个copy，要修改参数需要传递引用
	a := 10
	var a1 int
	a1 = setInput(a)
	fmt.Println(a)
	fmt.Println(a1)
	a2 := setInput2(&a)
	fmt.Println(a)
	fmt.Println(a2)
}

func setInput(a int) int {
	a = 100
	return a
}

func setInput2(a *int) int {
	*a = 100
	return *a
}
