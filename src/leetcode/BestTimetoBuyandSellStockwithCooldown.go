package leetcode

import "math"

/**
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
*/

/**
还是股票问题，加了一个冷冻期

dp[i][k][1] = max(dp[i-1][k][1], dp[i-2][k-1][0] - prices[i])
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
由于不限制交易次数，可以消除k
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
i状态之和i-1、i-2有关，还可以进行简化
*/
func maxProfitWithCool(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dpi0, dpi1, dpi2 := 0, math.MinInt32, 0
	for i := 0; i < len(prices); i++ {
		tmp := dpi0
		dpi1 = max(dpi1, dpi2-prices[i])
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi2 = tmp
	}
	return dpi0
}
