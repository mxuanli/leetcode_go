/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
*/

package main

import "fmt"

func reverseString(s []byte) []byte {
	i, j := 0, len(s)-1
	reverseR(i, j, &s)
	return s
}

func reverseR(i int, j int, s *[]byte) {
	if i >= j {
		return
	}
	reverseR(i+1, j-1, s)
	foo(i, j, s)
}

func foo(i int, j int, s *[]byte) {
	a := *s
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
	fmt.Println(a)
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	r := reverseString(s)
	fmt.Println(r)
}
