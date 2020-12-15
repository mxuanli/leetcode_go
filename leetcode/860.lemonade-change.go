/*

在柠檬水摊上，每一杯柠檬水的售价为 5 美元。

顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。

每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。

注意，一开始你手头没有任何零钱。

如果你能给每位顾客正确找零，返回 true ，否则返回 false 。

*/

package main

import "fmt"

func lemonadeChange(bills []int) bool {
	bills5 := 0
	bills10 := 0
	for _, b := range bills {
		if b == 5 {
			bills5++
		} else if b == 10 {
			if bills5 >= 1 { // 10块找五块
				bills5--
			} else {
				return false
			}
			bills10++
		} else if b == 20 {
			if bills10 >= 1 && bills5 >= 1 { // 20块优先找10块的
				bills10--
				bills5--
			} else if bills5 >= 3 {
				bills5 = bills5 - 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	bills := []int{5, 5, 5, 10, 20}
	r := lemonadeChange(bills)
	fmt.Println(r)
}
