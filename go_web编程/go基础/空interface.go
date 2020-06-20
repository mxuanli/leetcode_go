package main

import "fmt"

func main() {
	var a interface{}
	var i int = 5
	a = i
	fmt.Println(a)
	funcI1(i)
}

func funcI1(a interface{}) {
	fmt.Println(a)
}
