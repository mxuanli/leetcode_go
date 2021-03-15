package main

import "fmt"

func calculate(s string) int {
	r := 0
	sign := '+'
	s = s + "a"
	num := 0
	n := len(s)
	stack := make([]int, n)
	for _, i := range s {
		if i >= '0' && i <= '9' {
			num *= 10
			num += int(i - '0')
		} else if i == ' ' {
			continue
		} else {
			if sign == '+' {
				stack = append(stack, num)
			} else if sign == '-' {
				stack = append(stack, 0-num)
			} else if sign == '*' {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp*num)
			} else if sign == '/' {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp/num)
			}
			num = 0
			sign = i
		}
	}
	for _, i := range stack {
		r += i
	}
	return r
}

func main() {
	s := " 3+5 / 2 "
	r := calculate(s)
	fmt.Println(r)
}
