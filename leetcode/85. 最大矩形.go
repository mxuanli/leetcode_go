package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	width := make([][]int, len(matrix))
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		matrixI := matrix[i]
		width[i] = make([]int, len(matrixI))
		for j := 0; j < len(matrixI); j++ {
			if matrixI[j] == '1' && j == 0 {
				width[i][j] = 1
			} else if matrixI[j] == '1' && j > 0 {
				width[i][j] = width[i][j-1] + 1
			}
			minWidth := width[i][j]
			for k := i; k >= 0; k-- {
				height := i - k + 1
				if minWidth > width[k][j] {
					minWidth = width[k][j]
				}
				area := height * minWidth
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(matrix)
	matrix1 := [][]byte{{'1'}}
	r := maximalRectangle(matrix1)
	fmt.Println(r)
}
