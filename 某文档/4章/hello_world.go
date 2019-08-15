package main

import fm "fmt"

func main() {
	var n string = "hello world"
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fm.Print(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, "\n")
	print(day)
	print_(n)
}

func print_(n string) {
	fm.Print(n)
}
