package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println(a)
	fmt.Println(&a)
	var ptr1 *string = &a
	fmt.Println(ptr1)
	*ptr1 = "world"
	fmt.Println(a)
}
