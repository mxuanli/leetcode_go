package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	father := make(map[[2]string][2]string)
	accountMap := make(map[[2]string][][2]string)
	var find func(x [2]string) [2]string
	find = func(x [2]string) [2]string {
		root := x
		// 查找
		for father[root] != root {
			root = father[root]
		}
		// 压缩
		if root != x {
			fatherRoot := father[x]
			father[x] = root
			x = fatherRoot
		}
		return root
	}
	var merge func(x [2]string, y [2]string)
	merge = func(x [2]string, y [2]string) {
		xf := find(x)
		yf := find(y)
		if xf != yf {
			father[xf] = yf
			accountMap[yf] = append(accountMap[yf], accountMap[xf]...)
			delete(accountMap, xf)
		}
	}
	var add func(x [2]string)
	add = func(x [2]string) {
		_, bl := father[x]
		if bl == false {
			father[x] = x
			accountMap[x] = [][2]string{x}
		}
	}
	for _, account := range accounts {
		name, rootEmail := account[0], account[1]
		add([2]string{name, rootEmail})
		account := account[2:]
		for _, email := range account {
			add([2]string{name, email})
			// 如果两个账户都有共同的邮箱地址，在合并的时候，他们的图会被合并到一起
			merge([2]string{name, rootEmail}, [2]string{name, email})
		}
	}
	var res [][]string
	for key := range father {
		v, _ := father[key]
		if v == key {
			userName := accountMap[key][0][0]
			userAccount := []string{userName}
			for _, account := range accountMap[key] {
				userAccount = append(userAccount, account[1])
			}
			sort.Strings(userAccount)
			res = append(res, userAccount)
		}
	}
	return res
}

func main() {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	r := accountsMerge(accounts)
	fmt.Println(r)
}
