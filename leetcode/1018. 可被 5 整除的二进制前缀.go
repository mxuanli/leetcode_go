package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	r := 0
	var rSlice []bool
	for _, a := range A {
		r = (r<<1 + a) % 10
		rSlice = append(rSlice, r%5 == 0)
	}
	return rSlice
}

func main() {
	A := []int{0, 1, 1}
	r := prefixesDivBy5(A)
	fmt.Println(r)
}
