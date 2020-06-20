package main

import "fmt"

func main() {
	funcInt()
	funcString()
}

func funcInt() {
	// int类型
	var int1 = int(5)
	fmt.Println(int1)

	var int2 int = 10
	fmt.Println(int2)

	var int3 = 20
	fmt.Println(int3)
}

func funcString() {
	// 字符串
	var string1 = string("hello")
	fmt.Println(string1)

	var string2 string = "world"
	fmt.Println(string2)

	string3 := "xuanli"
	fmt.Println(string3)
}
