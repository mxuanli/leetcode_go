/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2
*/

package main

import "fmt"

func firstUniqChar(s string) int {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	r := firstUniqChar(s)
	fmt.Println(r)
}
