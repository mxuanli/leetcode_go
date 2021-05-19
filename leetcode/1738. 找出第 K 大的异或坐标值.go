package main

import (
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	matrixXOR := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		matrixN := make([]int, n+1)
		matrixXOR[i] = matrixN
	}
	res := []int{}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			matrixXOR[i][j] = matrixXOR[i][j-1] ^ matrixXOR[i-1][j] ^ matrixXOR[i-1][j-1] ^ matrix[i-1][j-1]
			res = append(res, matrixXOR[i][j])
		}
	}
	sort.Ints(res)
	return res[len(res)-k]
}
