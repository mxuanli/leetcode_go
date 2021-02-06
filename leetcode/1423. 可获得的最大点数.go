package main

import "fmt"

func sum(Slice []int) int {
	r := 0
	for _, s := range Slice {
		r = r + s
	}
	return r
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	start, end := 0, n-k
	sumElse := sum(cardPoints[start:end])
	minElse := sumElse
	for ; end < n; end++ {
		sumElse = sumElse - cardPoints[start] + cardPoints[end]
		if sumElse < minElse {
			minElse = sumElse
		}
		start++
	}
	r := sum(cardPoints) - minElse
	return r
}

func main() {
	cardPoints := []int{9, 7, 7, 9, 7, 7, 9}
	k := 7
	r := maxScore(cardPoints, k)
	fmt.Println(r)
}
