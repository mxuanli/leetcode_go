/*
290. 单词规律
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
*/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	sSlice := strings.Fields(s)
	if len(sSlice) != len(pattern) {
		return false
	}
	// 双hash表
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		pV, pOk := pMap[pattern[i]]
		sV, sOk := sMap[sSlice[i]]
		if pOk && pV != sSlice[i] || sOk && sV != pattern[i] {
			return false
		} else {
			pMap[pattern[i]] = sSlice[i]
			sMap[sSlice[i]] = pattern[i]
		}
	}
	return true
}

func wordPattern2(pattern string, s string) bool {
	sSlice := strings.Fields(s)
	if len(sSlice) != len(pattern) {
		return false
	}
	// 双hash表
	pMap := make(map[byte]string) // [pattern]s格式
	pVMap := make(map[string]int) // 存储value用
	for i := 0; i < len(pattern); i++ {
		pV, pOk := pMap[pattern[i]]
		if pOk {
			if pV != sSlice[i] { // 判断字符是否和前边同一个字符对应的值相等
				return false
			}
		} else {
			_, pOk1 := pVMap[sSlice[i]] // 判断value是否在之前被添加过
			if pOk1 {
				return false
			}
			pMap[pattern[i]] = sSlice[i]
			pVMap[sSlice[i]] = 1
		}
	}
	return true
}

func main() {
	pattern := "abba"
	str := "dog dog dog dog"
	r := wordPattern2(pattern, str)
	fmt.Println(r)
}
