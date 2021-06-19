package main

func maxLength(arr []string) int {
	var strSlice []string
	for _, a := range arr {
		if validStr(a) {
			continue
		}
		for _, s := range strSlice {
			if validStr(s+a) == false {
				strSlice = append(strSlice, s+a)
			}
		}
		strSlice = append(strSlice, a)
	}
	r := 0
	for _, i := range strSlice {
		if r < len(i) {
			r = len(i)
		}
	}
	return r
}

func validStr(text string) bool {
	textMap := make(map[rune]int)
	for _, i := range text {
		if textMap[i] > 0 {
			return true
		}
		textMap[i]++
	}
	return false
}
