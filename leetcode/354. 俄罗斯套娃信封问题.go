package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		envelopeI := envelopes[i]
		envelopeJ := envelopes[j]
		if envelopeI[0] < envelopeJ[0] {
			return true
		}
		if envelopeI[0] == envelopeJ[0] {
			if envelopeI[1] > envelopeJ[1] {
				return true
			}
		}
		return false
	})
	var newEnvelopes []int
	for _, v := range envelopes {
		_, v1 := v[0], v[1]
		newEnvelopes = append(newEnvelopes, v1)
	}
	db := make([]int, n)
	for i := 0; i < n; i++ {
		db[i] = 1
		for j := 0; j < i; j++ {
			if newEnvelopes[i] > newEnvelopes[j] {
				if db[j]+1 > db[i] {
					db[i] = db[j] + 1
				}
			}
		}
	}
	r := 0
	for _, v := range db {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	r := maxEnvelopes(envelopes)
	fmt.Println(r)
}
