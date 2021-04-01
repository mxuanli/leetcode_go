package main

func clumsy(N int) int {
	i := 0
	key := 2
	stack := []int{}
	for j := N; j > 0; j-- {
		if key == 0 {
			num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, num*j)
		} else if key == 1 {
			num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, num/j)
		} else if key == 2 {
			stack = append(stack, j)
		} else {
			stack = append(stack, 0-j)
		}
		key = i % 4
		i++
	}
	r := 0
	for _, v := range stack {
		r += v
	}
	return r
}
