package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			A[i][j], A[i][m-j-1] = A[i][m-j-1], A[i][j]
		}
		for j := 0; j < m; j++ {
			A[i][j] = A[i][j] ^ 1
		}
	}
	return A
}

func main() {
	A := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	r := flipAndInvertImage(A)
	fmt.Println(r)
}
