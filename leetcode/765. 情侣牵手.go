package main

import "fmt"

func minSwapsCouples(row []int) int {
	r, n := 0, len(row)
	for i := 0; i < n; i += 2 {
		if row[i]/2 == row[i+1]/2 { // 如果相邻的是情侣则跳过
			continue
		}
		for j := i + 1; j < n; j++ { // 如果不是情侣则找到情侣并交换
			if row[i]/2 == row[j]/2 {
				row[i+1], row[j] = row[j], row[i+1]
				r += 1
			}
		}
	}
	return r
}

func main() {
	row := []int{0, 2, 1, 3}
	r := minSwapsCouples(row)
	fmt.Println(r)
}
