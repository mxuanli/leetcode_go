package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println(len(str1))
	fmt.Println(len(str2))
	fmt.Println(utf8.RuneCountInString(str1))
	fmt.Println(utf8.RuneCountInString(str2))

}
