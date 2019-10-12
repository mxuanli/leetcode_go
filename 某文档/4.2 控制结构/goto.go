package main

import "fmt"

func main() {
	//gotoFunc()
	problem()
}

func gotoFunc() {
TestGoto:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 8 {
			fmt.Println("i为8，触发goto语句")
			goto TestGoto
		}
	}
}

func problem() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}
