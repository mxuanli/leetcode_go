package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sSlice := make([]int, len(s))
	tSlice := make([]int, len(t))
	sMap := make(map[int32]int)
	tMap := make(map[int32]int)
	for index, v := range s {
		iV := sMap[v]
		if iV == 0 {
			sMap[v] = index + 1
			sSlice = append(sSlice, index+1)
		} else {
			sSlice = append(sSlice, iV)
		}
	}
	for index, v := range t {
		iV := tMap[v]
		if iV == 0 {
			tMap[v] = index + 1
			tSlice = append(tSlice, index+1)
		} else {
			tSlice = append(tSlice, iV)
		}
	}
	if len(sSlice) != len(tSlice) {
		return false
	}
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] != tSlice[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "egg"
	t := "ada"
	r := isIsomorphic(s, t)
	fmt.Println(r)
}
