package main

import "fmt"

//
//type NumArray struct {
//	nums []int
//}
//
//
//func Constructor(nums []int) NumArray {
//	var n NumArray
//	n.nums = nums
//	return n
//}
//
//func sum(Slice []int) int {
//	r := 0
//	for _, s := range Slice {
//		r = r + s
//	}
//	return r
//}
//
//func (this *NumArray) SumRange(i int, j int) int {
//	r := sum(this.nums[i: j + 1])
//	return r
//}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	var n NumArray
	for i := 0; i <= len(nums); i++ {
		n.nums = append(n.nums, 0)
	}
	for i, v := range nums {
		n.nums[i+1] = n.nums[i] + v
	}
	return n
}

func (this *NumArray) SumRange(i int, j int) int {
	r := this.nums[j+1] - this.nums[i]
	return r
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	n := Constructor(nums)
	r := n.SumRange(0, 2)
	fmt.Println(r)
	r = n.SumRange(2, 5)
	fmt.Println(r)
	r = n.SumRange(0, 5)
	fmt.Println(r)
}
