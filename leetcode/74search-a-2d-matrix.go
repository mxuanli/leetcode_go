package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	start := 0
	len_matrix := len(matrix[0])
	end := len_matrix*len(matrix) - 1
	if end < 0 {
		return false
	}
	if matrix[0][0] == target {
		return true
	}
	if matrix[len(matrix)-1][len_matrix-1] == target {
		return true
	}
	for start+1 < end {
		mid := (start + end) / 2
		mid_m := matrix[mid/len_matrix][mid%len_matrix]
		if mid_m == target {
			return true
		} else if mid_m < target {
			start = mid
		} else if mid_m > target {
			end = mid
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	var matrix1 [][]int
	for i := 0; i < len(matrix)-1; i++ {
		append(matrix)
	}
}

func main() {
	matrix_1 := []int{1, 3, 5, 7}
	matrix_2 := []int{10, 11, 16, 20}
	matrix_3 := []int{23, 30, 34, 50}
	matrix := [][]int{matrix_1, matrix_2, matrix_3}
	//matrix := [][]int{matrix_1}
	//matrix := [][]int{}
	target := 3
	r := searchMatrix2(matrix, target)
	fmt.Println(r)
}
