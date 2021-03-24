package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	nums := make([]int, len(tokens))
	for _, token := range tokens {
		if token != "+" && token != "-" && token != "/" && token != "*" {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		} else {
			num1 := nums[len(nums)-1]
			num2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			num3 := 0
			switch token {
			case "+":
				num3 = num2 + num1
			case "-":
				num3 = num2 - num1
			case "/":
				num3 = int(num2 / num1)
			case "*":
				num3 = num2 * num1
			}
			nums = append(nums, num3)
		}
	}
	return nums[len(nums)-1]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	r := evalRPN(tokens)
	fmt.Println(r)
}
