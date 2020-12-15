/*
给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。

（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）


示例 3:

输入: N = 332
输出: 299
*/

package main

import "fmt"

func monotoneIncreasingDigits(N int) int {
	// 转换为slice
	var sliceN []int
	for N != 0 {
		tmp := N % 10
		N = N / 10
		sliceN = append(sliceN, tmp)
	}
	for i, j := 0, len(sliceN)-1; i < j; i, j = i+1, j-1 {
		sliceN[i], sliceN[j] = sliceN[j], sliceN[i]
	}
	for i := 0; i < len(sliceN)-1; i++ {
		if sliceN[i] > sliceN[i+1] { // 如果左边的比右边的大
			sliceN[i]--                            // 就把左边的减1
			for j := i + 1; j < len(sliceN); j++ { // 右边的都设置为9
				sliceN[j] = 9
			}
			i = -1 // 设置为-1，从头遍历
		}
	}
	r := 0
	for _, v := range sliceN {
		r = r * 10
		r = r + v
	}
	return r
}

func main() {
	N := 100
	r := monotoneIncreasingDigits(N)
	fmt.Println(r)
}
