package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r[j][i] = matrix[i][j]
		}
	}
	return r
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := transpose(matrix)
	fmt.Println(r)
}
