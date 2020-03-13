package main

import "fmt"

const Pi float64 = 3.14

const (
	i = 10
	j = "20"
	k = 10.1
)

const (
	a       = iota // a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func main() {
	fmt.Println(Pi)
	fmt.Println(i)
}
