package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	r := 0
	lenF := len(flowerbed)
	for i := 0; i < lenF; i++ {
		front, back := 1, 1
		if flowerbed[i] == 0 {
			if i == 0 {
				front = 0
			} else {
				front = flowerbed[i-1]
			}
			if i == lenF-1 {
				back = 0
			} else {
				back = flowerbed[i+1]
			}
			if front == 0 && back == 0 {
				flowerbed[i] = 1
				r++
			}
		}
	}
	if r >= n {
		return true
	}
	return false
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2
	r := canPlaceFlowers(flowerbed, n)
	fmt.Println(r)
}
