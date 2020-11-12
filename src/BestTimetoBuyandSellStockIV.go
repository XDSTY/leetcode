package main

import (
	"math"
)

/**
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
Design an algorithm to find the maximum profit. You may complete at most k transactions.
Notice that you may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
还是股票问题，限制最多只能进行k此交易

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
*/

/**
还是假设买入股票就认为完成一次交易，即买入股票k就加一
 dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
 dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
*/

func main() {
	maxProfitWithK(2, []int{3, 2, 6, 5, 0, 3})
}

func maxProfitWithK(k int, prices []int) int {
	// 大于 prices.length/2，表明不限制交易的次数，可以复用 股票|| 的代码
	if k > len(prices)/2 {
		return maxProfitIIWithNolimit(prices)
	}
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[0][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][k][0]
}

func maxProfitIIWithNolimit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp0, dp1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		tmp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, tmp-prices[i])
	}
	return dp0
}

func max(val1, val2 int) int {
	if val1 < val2 {
		return val2
	}
	return val1
}
