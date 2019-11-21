package main

// map 就是 字典

import (
	"fmt"
)

func main() {
	var mapA map[string]string     // 声明字典A
	mapA = make(map[string]string) // 使用之前使用make初始化
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
	mapThree, ok := mapB["three"]
	if ok {
		fmt.Println(mapThree)
	}
	Map()
}

func Map() {
	var mapB map[string]string
	mapB = make(map[string]string) // 初始化
	mapB["one"] = "1"
	mapB["two"] = "2"
	mapB["three"] = "3"
	fmt.Println(mapB["one"], mapB["two"])
	var k string
	var j string
	for k, j = range mapB {
		fmt.Println(k, j)
	}
}
