package main

import (
	"fmt"
)

func main() {
	var mapA map[string]string // 声明字典A
	mapA = make(map[string]string)
	mapA["one"] = "1"
	mapA["two"] = "2"
	mapA["three"] = "3"
	fmt.Println("第一个值是", mapA["one"])
	mapB := make(map[string]int) // 声明字典B
	mapB["one"] = 1
	mapB["two"] = 2
	mapB["three"] = 3
	fmt.Println("第二个元素", mapB["two"])
	for k, v := range mapB {
		fmt.Println(k, v)
	}
}
