package main

import "fmt"

func sum(Slice []int) int {
	r := 0
	for _, s := range Slice {
		r = r + s
	}
	return r
}

func fairCandySwap(A []int, B []int) []int {
	aver := (sum(A) - sum(B)) / 2
	var mapA map[int]int
	mapA = make(map[int]int)
	for _, j := range A {
		mapA[j] = j
	}
	for _, i := range B {
		j := i + aver
		_, ok := mapA[j]
		if ok {
			return []int{j, i}
		}

	}
	return []int{}
}

func main() {
	A := []int{1, 1}
	B := []int{2, 2}
	r := fairCandySwap(A, B)
	fmt.Println(r)
}
