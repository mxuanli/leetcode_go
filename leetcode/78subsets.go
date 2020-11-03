package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int // 储存结果
	var list []int     // 储存临时结果
	backtrack(0, list, &result, nums)
	return result
}

func backtrack(i int, list []int, result *[][]int, nums []int) {
	ans := make([]int, len(list)) // 拷贝临时结果，避免受到影响
	copy(ans, list)
	*result = append(*result, ans)
	for j := i; j < len(nums); j++ { // 基线条件，当j等于num的长度时，结束一次递归
		list = append(list, nums[j]) // 添加一个元素到集合内
		backtrack(j+1, list, result, nums)
		list = list[0 : len(list)-1] // 回撤，尝试其它路径
	}
}

func main() {
	nums := []int{9, 0, 3, 5, 7}
	c := subsets(nums)
	fmt.Println(c)
}
