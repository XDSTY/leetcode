package leetcode

import "math"

/**
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
返回获得利润的最大值。
注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
*/

/**
还是股票交易，增加了手续费
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i] - fee)
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
k是不限次数的
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
i之和i-1有关
*/
func maxProfitWithTransactionFee(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}
	dpi1, dpi0 := math.MinInt32, 0
	for i := 0; i < len(prices); i++ {
		dpi1 = max(dpi1, dpi0-prices[i]-fee)
		dpi0 = max(dpi0, dpi1+prices[i])
	}
	return dpi0
}
