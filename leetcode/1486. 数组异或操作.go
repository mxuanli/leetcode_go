package main

func xorOperation(n int, start int) int {
	r := 0
	for i := 0; i < n; i++ {
		num := start + 2*i
		r ^= num
	}
	return r
}
