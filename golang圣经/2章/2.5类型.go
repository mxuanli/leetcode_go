package main

import fm "fmt"

type TestType1 string
type TestType2 string

func main() {
	var a TestType1
	a = "hello"
	fm.Println(a)
	var b TestType2
	b = "hello"
	fm.Println(b)
	//fm.Println(a == b)  //两个类型不通的变量无法比较
	fm.Println(a == TestType1(b)) //转换类型之后可以比较
}
