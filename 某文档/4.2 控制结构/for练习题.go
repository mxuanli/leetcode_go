package main

import "fmt"

func main() {
	//printG()
	//print35()
	test()
}

func test() {

}

func printG() {
	//生成
	//G
	//GG
	//GGG
	//GGGG
	//GGGGG
	//GGGGGG
	for i := 1; i <= 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}

func print35() {
	for i := 1; i <= 100; i++ {
		switch true {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}
	}

}
