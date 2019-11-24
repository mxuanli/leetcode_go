package main

import "fmt"

func main() {
	throwsPanic(funcF)
}

func funcF() {
	fmt.Println("执行1")
	panic("发生错误")
	fmt.Println("执行2")
}

func throwsPanic(f func()) (b string) {
	defer func() {
		//if x := recover(); x != nil {
		//	b = "错误已恢复"
		//	fmt.Println(b)
		//}
		x := recover()
		if x != nil {
			fmt.Println(x)
			fmt.Println("错误已恢复")
		}
	}()
	f() // 执行函数f，如果f中出现了panic，那么就可以恢复回来
	fmt.Println("正常执行")
	return
}
