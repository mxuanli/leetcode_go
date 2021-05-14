package main

import "fmt"

func intToRoman(num int) string {
	numHash := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanHash := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var r string
	for i, v := range numHash {
		if num/v > 0 {
			romanCount := num / v
			for c := 0; c < romanCount; c++ {
				r += romanHash[i]
			}
			num %= v
		}
	}
	return r
}

func main() {
	num := 3998
	r := intToRoman(num)
	fmt.Println(r)
}
