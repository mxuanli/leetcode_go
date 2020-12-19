/*
389. 找不同
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。


示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
*/

package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var r uint8
	r = 0
	for si := 0; si < len(s); si++ {
		r = r ^ s[si]
	}
	for ti := 0; ti < len(t); ti++ {
		r = r ^ t[ti]
	}
	return byte(r)
}

func findTheDifference1(s string, t string) byte {
	// 转unicode，相加再相减，再在转回字符串
	sc, tc := 0, 0
	for si := 0; si < len(s); si++ {
		sc = sc + int(s[si])
	}
	for ti := 0; ti < len(t); ti++ {
		tc = tc + int(t[ti])
	}
	return byte(tc - sc)
}

func main() {
	s := "abcd"
	t := "abcde"
	r := findTheDifference1(s, t)
	fmt.Println(r)
}
