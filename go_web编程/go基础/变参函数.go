package main

import (
	"fmt"
	"reflect"
)

func main() {
	moreArg(1, 2, 3, 4)
}

func moreArg(arg ...int) {
	fmt.Println(arg)
	fmt.Println(reflect.TypeOf(arg))
	for k, v := range arg {
		fmt.Println(k, v)
	}
}
