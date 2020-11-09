package main

import "fmt"

func reverseString(s []byte) []byte {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

func reverseString2(s []byte) {
	i, j := 0, len(s)-1
	reverseR(i, j, &s)
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
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o', 'w'}
	fmt.Println(s)
	r := reverseString(s)
	fmt.Println(r)
}
