package main

import fm "fmt"

func main() {
	x := new(int)
	fm.Println(x)
	fm.Println(*x)
	*x = 10
	fm.Println(*x)
}
