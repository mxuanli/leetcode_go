package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	xors[0] = 0
	for i, a := range arr {
		xors[i+1] = xors[i] ^ a
	}
	r := make([]int, len(queries))
	for i, q := range queries {
		q1 := q[0]
		q2 := q[1]
		r[i] = xors[q1] ^ xors[q2+1]
	}
	return r
}

func main() {
	arr := []int{16}
	queries := [][]int{{0, 0}, {0, 0}, {0, 0}}
	r := xorQueries(arr, queries)
	fmt.Println(r)
}
