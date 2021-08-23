package main

func isValid(s string) bool {
	dictS := make(map[byte]byte)
	dictS['('] = ')'
	dictS['{'] = '}'
	dictS['['] = ']'
	var stack []byte
	for j := range s {
		i := s[j]
		// 如果是前括号
		if dictS[i] != 0 {
			stack = append(stack, dictS[i])
		} else {
			// 如果是后括号
			if len(stack) == 0 || stack[len(stack)-1] != i {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
