package main

import "fmt"

func addToArrayForm2(A []int, K int) []int {
	var KSlice []int
	// 转换为slice
	for K != 0 {
		k := K % 10
		KSlice = append(KSlice, k)
		K = K / 10
	}
	// 相加
	An := len(A) - 1
	Kn := len(KSlice)
	tmp := 0
	Ai := An
	for i := 0; i < Kn; i++ {
		Ai = An - i
		if Ai < 0 {
			Ain := 0
			Kin := KSlice[i]
			AKi := Ain + Kin + tmp
			tmp = AKi / 10
			AKi = AKi % 10
			A = append([]int{AKi}, A...)
		} else {
			Ain := A[Ai]
			Kin := KSlice[i]
			AKi := Ain + Kin + tmp
			tmp = AKi / 10
			AKi = AKi % 10
			A[Ai] = AKi
		}

	}
	// 进位
	for Ai != 0 && tmp != 0 {
		Ai -= 1
		if Ai >= 0 {
			Ain := A[Ai]
			AKi := Ain + tmp
			tmp = AKi / 10
			AKi = AKi % 10
			A[Ai] = AKi
		} else {
			Ain := 0
			AKi := Ain + tmp
			tmp = AKi / 10
			AKi = AKi % 10
			A = append([]int{AKi}, A...)
		}
	}
	if tmp != 0 {
		A = append([]int{tmp}, A...)
	}
	return A
}

func addToArrayForm(A []int, K int) []int {
	var KSlice []int
	// 转换为slice
	for K != 0 {
		k := K % 10
		KSlice = append(KSlice, k)
		K = K / 10
	}
	// 相加
	An := len(A) - 1
	Kn := len(KSlice)
	Kin := 0
	tmp := 0
	for i := An; i >= 0; i-- {
		Ain := A[i]
		Ki := An - i
		if Ki < Kn {
			Kin = KSlice[i]
		} else {
			Kin = 0
		}
	}
}

func main() {
	A := []int{7}
	K := 993
	r := addToArrayForm(A, K)
	fmt.Println(r)
}
