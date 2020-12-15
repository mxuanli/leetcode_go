/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

先定义一个map，key为string，值为slice
先把元素排序，排序后作为map的key，把元素添加到slice里，这样异位词会被添加到同一个slice里
遍历结束后，取出所有的value
*/

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		key := getKey(bytes)
		tmp[key] = append(tmp[key], str) // r[排序后][排序前1, 排序前2]
	}
	// 取出value
	var r [][]string
	for _, str := range tmp {
		r = append(r, str)
	}
	return r
}

func getKey(bytes []byte) string {
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	r := groupAnagrams(strs)
	fmt.Println(r)
}
