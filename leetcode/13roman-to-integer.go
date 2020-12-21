/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
*/

package main

import "fmt"

func romanToInt(s string) int {
	var a map[byte]int
	a = make(map[byte]int)
	a['I'] = 1
	a['V'] = 5
	a['X'] = 10
	a['L'] = 50
	a['C'] = 100
	a['D'] = 500
	a['M'] = 1000
	r := 0
	for i := 0; i < len(s)-1; i++ {
		// 左边小于右边，则减去左边，左边大于右边则相加左边
		if a[s[i]] < a[s[i+1]] {
			r = r - a[s[i]]
		} else {
			r = r + a[s[i]]
		}
	}
	r = r + a[s[len(s)-1]]
	return r
}

func main() {
	s := "III"
	r := romanToInt(s)
	fmt.Println(r)
}
