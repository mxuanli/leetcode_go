package main

import "fmt"

func setZeroes2(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	zeroN := make([]bool, n)
	zeroM := make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zeroN[i] = true
				zeroM[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if zeroN[i] || zeroM[j] {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	flag := false
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if flag {
			matrix[i][0] = 0
		}
	}
	fmt.Println(matrix)
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
}
