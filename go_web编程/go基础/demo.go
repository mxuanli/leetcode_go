package main

import "fmt"

func fibonacci(c chan int, n int) {
	i, j := 1, 1
	for k := 0; k < n; k++ {
		c <- i
		i, j = j, i+j
	}
	close(c)
}

func main() {
	a := 10
	var b *int
	fmt.Println(&a)
	b = &a
	fmt.Println(b)
	fmt.Println(*b)
}
