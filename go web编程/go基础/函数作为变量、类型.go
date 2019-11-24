package main

import (
	"fmt"
)

//定义了一个函数类型
type funcTestInt func(int) int

func main() {
	func1(func2)
}

//函数当作了参数传递
func func1(func3 funcTestInt) {
	a := 15
	func2Return := func3(a)
	fmt.Println(func2Return)
}

func func2(a int) int {
	b := a + 10
	return b
}
