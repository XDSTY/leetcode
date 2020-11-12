package main

import "math"

/**
Say you have an array prices for which the ith element is the price of a given stock on day i.
给你一个数组，数组中的，每个元素的下标代表着天数，值代表着这天的股票价格
Design an algorithm to find the maximum profit.
You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
设计一个算法找到最大的收益，可以进行多次的交易
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
不能同时进行交易

可以进行任意次交易

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
*/

func maxProfitII(prices []int) int {
	total := 0
	minEle := prices[0]
	prices = append(prices, 0)
	for i := 1; i < len(prices); i++ {
		// 进入降序
		if prices[i] < prices[i-1] {
			total += prices[i-1] - minEle
			minEle = prices[i]
		}
	}
	return total
}

/**
0 <= i < len(prices), 0 <= k <= len(prices) / 2
假设认为买了股票才是一笔交易的完成
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
因为k为正无穷，可以认为k和k-1是一样大的，则上述公式可以为：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])
可以进行简化
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
i只和i-1有关，则还可以进行简化
dp0 = max(dp0, dp1 + prices[i])
dp1 = max(dp1, dp0 - prices[i])
*/
func maxProfitIIWithDp(prices []int) int {
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
