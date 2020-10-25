package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	result := 0
	k := 1
	if len(s) == 0 {
		return result
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}
	if s[0] != '-' && s[0] != '+' && !(s[0] >= '0' && s[0] <= '9') {
		return result
	}
	for i, v := range s {
		if v == '-' || v == '+' {
			if i == 0 {
				if v == '-' {
					k = 0
				} else {
					k = 1
				}
				continue
			}
			break
		}
		if v >= '0' && v <= '9' {
			result = result * 10
			a := int(v - '0')
			result = result + a
		} else {
			break
		}
		if k > 0 && result > math.MaxInt32 {
			return math.MaxInt32
		} else if k == 0 && result > math.MaxInt32+1 {
			return 0 - math.MaxInt32 - 1
		}
	}
	if k == 0 {
		result = 0 - result
	}
	return result
}

func main() {
	s := "00000-42a1234"
	result := myAtoi(s)
	fmt.Println(result)
}
