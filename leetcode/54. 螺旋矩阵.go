package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var r []int
	n, m := len(matrix), len(matrix[0])
	top, right, down, left := 0, m-1, n-1, 0
	x, y := 0, 0
	for len(r) < n*m {
		for x == top && y <= right {
			value := matrix[x][y]
			r = append(r, value)
			y += 1
		}
		top += 1
		x, y = x+1, y-1
		for y == right && x <= down {
			value := matrix[x][y]
			r = append(r, value)
			x += 1
		}
		right -= 1
		x, y = x-1, y-1
		for top <= down && x == down && y >= left {
			value := matrix[x][y]
			r = append(r, value)
			y -= 1
		}
		down -= 1
		x, y = x-1, y+1
		for left <= right && y == left && x >= top {
			value := matrix[x][y]
			r = append(r, value)
			x -= 1
		}
		left += 1
		x, y = x+1, y+1
	}
	return r
}

func main() {
	matrix := [][]int{{1}, {4}, {7}}
	r := spiralOrder(matrix)
	fmt.Println(r)
}
