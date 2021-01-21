package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]
	diffX, diffY := x2-x1, y2-y1
	for _, coordinate := range coordinates[2:] {
		xi, yi := coordinate[0], coordinate[1]
		diffXi := xi - x1
		diffYi := yi - y1
		if (diffX*diffYi)-(diffXi*diffY) != 0 {
			return false
		}
	}
	return true
}

func main() {
	//coordinates := [][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}}
	coordinates := [][]int{{0, 0}, {0, 1}, {0, -1}}
	r := checkStraightLine(coordinates)
	fmt.Println(r)
}
