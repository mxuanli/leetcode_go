package main

import "fmt"

func abs(x int) int {
	// 计算绝对值
	if x >= 0 {
		return x
	}
	return 0 - x
}

func equalSubstring(s string, t string, maxCost int) int {
	// 计算转换需要的花费
	var costSlice []int
	n := len(s)
	for i := 0; i < n; i++ {
		cost := abs(int(s[i]) - int(t[i]))
		costSlice = append(costSlice, cost)
	}
	// 找到在最大花费范围内的最长窗口
	start, cost, r := 0, 0, 0
	for end := 0; end < n; end++ {
		cost += costSlice[end]
		for cost > maxCost {
			cost -= costSlice[start]
			start++
		}
		thisR := end - start + 1
		if r < thisR {
			r = thisR
		}
	}
	return r
}

func main() {
	s := "abcd"
	t := "bcdf"
	cost := 3
	r := equalSubstring(s, t, cost)
	fmt.Println(r)
}
