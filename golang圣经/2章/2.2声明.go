package main

import "fmt"

const pai = 3.14 // 常量
const boilingF = 212.0

func main() {
	var a float64 = pai
	r := 10
	var c = a * float64(r) * float64(r)
	fmt.Println(c)
}
