/*
330. 按要求补齐数组
给定一个已排序的正整数数组 nums，和一个正整数 n 。从 [1, n] 区间内选取任意个数字补充到 nums 中，使得 [1, n] 区间内的任何数字都可以用 nums 中某几个数字的和来表示。请输出满足上述要求的最少需要补充的数字个数。

示例 1:

输入: nums = [1,3], n = 6
输出: 1
解释:
根据 nums 里现有的组合 [1], [3], [1,3]，可以得出 1, 3, 4。
现在如果我们将 2 添加到 nums 中， 组合变为: [1], [2], [3], [1,3], [2,3], [1,2,3]。
其和可以表示数字 1, 2, 3, 4, 5, 6，能够覆盖 [1, 6] 区间里所有的数。
所以我们最少需要添加一个数字。
*/

package main

import "fmt"

func minPatches(nums []int, n int) int {
	r := 0
	x := 1
	lenNums := len(nums)
	index := 0
	for x <= n {
		// x前边的值都已经被覆盖，如果nums[index]不大于x，则x也会被覆盖
		if index < lenNums && nums[index] <= x {
			// 如果x被覆盖，那么x + nums[index] - 1也会被覆盖
			x = x + nums[index]
			index = index + 1
			// 如果nums[index]大于x，则往里加数
		} else {
			x = x * 2
			r = r + 1
		}
	}
	return r
}

func main() {
	nums := []int{1, 3}
	n := 6
	r := minPatches(nums, n)
	fmt.Println(r)
}
