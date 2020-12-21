/*
746. 使用最小花费爬楼梯
数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:

输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
*/
package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	/*
		每一步需要消耗的体力为前一步或者前两步较小的值加上本步的值，cost[i] + min(r[i - 1], r[i - 2])
		存储到每一步的最小值，最后从倒数两个选取最小的
	*/
	n := len(cost)
	var r [1000]int
	r[0], r[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		if r[i-1] < r[i-2] {
			r[i] = cost[i] + r[i-1]
		} else {
			r[i] = cost[i] + r[i-2]
		}
	}
	result := 0
	if r[n-1] < r[n-2] {
		result = r[n-1]
	} else {
		result = r[n-2]
	}
	return result
}

func main() {
	cost := []int{1, 1, 0, 0}
	r := minCostClimbingStairs(cost)
	fmt.Println(r)
}
