package main

import "fmt"

func isValidSerialization(preorder string) bool {
	stack := []int{1}
	for i := 0; i < len(preorder); i++ {
		pre := preorder[i]
		if len(stack) == 0 {
			return false
		}
		if pre == ',' {
			continue
		} else if pre >= '0' && pre <= '9' {
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 2)
		} else if pre == '#' {
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	preorder := "9,#,92,#,#"
	r := isValidSerialization(preorder)
	fmt.Println(r)
}
