/*
135. 分发糖果
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
*/
package main

import (
	"fmt"
)

func candy(ratings []int) int {
	// 贪心
	lenRatings := len(ratings)
	// 初始值都为1
	leftArray := make([]int, lenRatings)
	rightArray := make([]int, lenRatings)
	for i := 0; i < lenRatings; i++ {
		leftArray[i] = 1
		rightArray[i] = 1
	}
	// 左边取最优
	for i := 1; i < lenRatings; i++ {
		if ratings[i] > ratings[i-1] {
			leftArray[i] = leftArray[i-1] + 1
		}
	}
	// 右边取最优
	for i := lenRatings - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rightArray[i] = rightArray[i+1] + 1
		}
	}
	// 合并取最大，就是既符合左边又符合右边
	r := 0
	for i := 0; i < lenRatings; i++ {
		if leftArray[i] > rightArray[i] {
			r = r + leftArray[i]
		} else {
			r = r + rightArray[i]
		}
	}
	return r
}

func main() {
	ratings := []int{1, 2, 87, 87, 87, 2, 1}
	r := candy(ratings)
	fmt.Println(r)
}
