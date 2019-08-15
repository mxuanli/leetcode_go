package main

import "fmt"

func main() {
	a := "hel" + "lo"
	b := a + "world"
	fmt.Println(b)
	fmt.Println(a[1])
	fmt.Println(len(b))
}
