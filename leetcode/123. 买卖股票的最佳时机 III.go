package main

import "fmt"

func maxProfit(prices []int) int {
	// 动态规划
	lenP := len(prices)
	buy1, buy2 := -prices[0], -prices[0]
	sell1, sell2 := 0, 0
	var max func(int, int) int
	max = func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < lenP; i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	r := maxProfit(prices)
	fmt.Println(r)
}
