/*
714. 买卖股票的最佳时机含手续费
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

示例 1:

输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
输出: 8
解释: 能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
*/

package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	buy, sell := prices[0], prices[0]
	r := 0
	for i := 0; i < len(prices); i++ {
		if buy < sell-fee {
			if prices[i] < sell-fee { // 卖出时得到的钱，大于当日的买入价格，就卖了重新买
				r = r + sell - fee - buy
				buy, sell = prices[i], prices[i]
			} else if prices[i] > sell {
				sell = prices[i]
			}
		} else {
			if prices[i] > sell {
				sell = prices[i]
			}
			if prices[i] < buy {
				buy, sell = prices[i], prices[i]
			}
		}
	}
	if buy < sell-fee {
		r = r + sell - fee - buy
	}
	return r
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	r := maxProfit(prices, fee)
	fmt.Println(r)
}
