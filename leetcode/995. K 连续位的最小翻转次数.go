package main

import "fmt"

func minKBitFlips(A []int, K int) int {
	n := len(A)
	var que []int
	r := 0
	for i := 0; i < n; i++ {
		if len(que) > 0 && i >= que[0]+K {
			que = que[1:]
		}
		if len(que)%2 == A[i] {
			if i+K > n {
				return -1
			}
			r++
			que = append(que, i)
		}
	}
	return r
}

func main() {
	A := []int{0, 1, 0}
	K := 1
	r := minKBitFlips(A, K)
	fmt.Println(r)
}
