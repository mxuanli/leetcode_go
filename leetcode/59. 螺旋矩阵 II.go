package main

import "fmt"

func generateMatrix(n int) [][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
	}
	var nSlice []int
	for i := 1; i <= n*n; i++ {
		nSlice = append(nSlice, i)
	}
	top, right, down, left := 0, n-1, n-1, 0
	x, y := 0, -1
	for len(nSlice) > 0 {
		for x == top && y < right {
			y += 1
			r[x][y] = nSlice[0]
			nSlice = nSlice[1:]
		}
		top += 1
		for y == right && x < down {
			x += 1
			r[x][y] = nSlice[0]
			nSlice = nSlice[1:]
		}
		right -= 1
		for x == down && y > left {
			y -= 1
			r[x][y] = nSlice[0]
			nSlice = nSlice[1:]
		}
		down -= 1
		for y == left && x > top {
			x -= 1
			r[x][y] = nSlice[0]
			nSlice = nSlice[1:]
		}
		left += 1
	}
	return r
}

func main() {
	n := 3
	r := generateMatrix(n)
	fmt.Println(r)
}
