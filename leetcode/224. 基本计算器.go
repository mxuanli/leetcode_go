package main

import "fmt"

func calculate(s string) int {
	tmp := 0
	r := 0
	sign := 1
	n := len(s)
	i := 0
	stack := make([]int, n)
	for i < n {
		si := s[i]
		if si == ' ' {
			i++
		} else if si == '+' {
			sign = 1
			i++
		} else if si == '-' {
			sign = -1
			i++
		} else if si == '(' {
			stack = append(stack, r)
			stack = append(stack, sign)
			sign = 1
			r = 0
			i++
		} else if si == ')' {
			sign = stack[len(stack)-1]
			tmp = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			r = tmp + r*sign
			i++
		} else {
			tmp = int(si - '0')
			i++
			for i < n && s[i] >= '0' && s[i] <= '9' {
				si = s[i]
				tmp *= 10
				tmp += int(si - '0')
				i++
				si = s[i]
			}
			tmp *= sign
			r += tmp
		}
	}
	return r
}

func main() {
	s := "2147483647"
	r := calculate(s)
	fmt.Println(r)
}
