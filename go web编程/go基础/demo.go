package main

import "fmt"

type months map[string]int

func main() {
	m := months{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	fmt.Println(m["1"])
}
