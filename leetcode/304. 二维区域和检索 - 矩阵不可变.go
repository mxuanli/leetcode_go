package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := NumMatrix{matrix}
	return n
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	r := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			r += this.matrix[i][j]
		}
	}
	return r
}

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	n := Constructor(matrix)
	r := n.SumRegion(2, 1, 4, 3)
	fmt.Println(r)
	r = n.SumRegion(1, 1, 2, 2)
	fmt.Println(r)
	r = n.SumRegion(1, 2, 2, 4)
	fmt.Println(r)
}
