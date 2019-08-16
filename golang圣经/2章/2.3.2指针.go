package main

import fm "fmt"

func main() {
	x := 1
	p := &x
	fm.Println("p指针指向的地址", p)
	fm.Println("p指针指向的值", *p)
	*p = 10
	fm.Println("修改后的x", x)
	var v1 = f()
	fm.Println(*v1)
}

func f() *int {
	v := 10
	return &v
}
