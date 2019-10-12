package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "sfhasffeofifvs"
	str2 := "fsjlglsgdfcgjsgdfdcdffs"
	int1 := 10
	float1 := 8.99
	var int_str_1 int
	fmt.Println(strings.HasPrefix(str1, "s"))
	fmt.Println(strings.HasSuffix(str1, "k"))
	fmt.Println(strings.HasSuffix(str1, "s"))
	fmt.Println(strings.Contains(str2, "d"))
	fmt.Println(strings.Contains(str2, "c"))
	fmt.Println(strings.Index(str2, "d"))
	fmt.Println(strings.LastIndex(str2, "d"))
	fmt.Println(strconv.Itoa(int1))
	fmt.Println(strconv.FormatFloat(float1, 'b', 10, 32))
	int_str_1, _ = strconv.Atoi(str1)
	fmt.Println(int_str_1)
	fmt.Println(strconv.IntSize)
}
