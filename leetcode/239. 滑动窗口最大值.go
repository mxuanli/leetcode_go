package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var (
		queue []int
		r     []int
	)
	for i := 0; i < len(nums); i++ {
		// 保持queue里第一个值是最大的
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 如果最大的值不在范围内，就删除了
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		// 如果满足滑动窗口的条件，就把最大值添加到结果里
		if i >= k-1 {
			r = append(r, nums[queue[0]])
		}
	}
	return r
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	r := maxSlidingWindow(nums, k)
	fmt.Println(r)
}
