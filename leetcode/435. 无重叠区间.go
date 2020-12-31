package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	lenIntervals := len(intervals)
	if lenIntervals == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else {
			return false
		}
	})
	r := 0
	for i, j := 0, 1; i < lenIntervals && j < lenIntervals; i, j = i+1, j+1 {
		// 如果左边的右区间大于右边的左区间，则发生重叠，需要移除区间
		if intervals[i][1] > intervals[j][0] {
			r++
			// 保留下右区间大的区间
			if intervals[i][1] < intervals[j][1] {
				i = i - 1 // 左边的小的话，把i-1，后边循环+=1的时候，i就不会变
			} else {
				i = j - 1 // 右边小的话，i = j - 1，后边循环+=1的时候，i会等于现在的j
			}
		} else {
			i = j - 1
		}
	}
	return r
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	r := eraseOverlapIntervals(intervals)
	fmt.Println(r)
}
