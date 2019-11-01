package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("测试错误")
	if err != nil {
		fmt.Println(err)
	}
}
