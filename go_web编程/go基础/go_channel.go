package main

import "fmt"

func Sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func RunSum() {
	a := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go Sum(a, c)
	LenA := len(a)
	fmt.Println(a[:LenA/2])
	go Sum(a[:LenA/2], c)
	s1, s2 := <-c, <-c
	fmt.Println(s1)
	fmt.Println(s2)
}

func RunChnn() {
	c := make(chan int, 2)
	c <- 2
	c <- 3
	c1 := <-c
	c2 := <-c
	fmt.Println(c1)
	fmt.Println(c2)
}

func main() {
	//RunSum()
	RunChnn()
}
