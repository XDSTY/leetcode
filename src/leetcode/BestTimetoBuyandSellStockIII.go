package leetcode

import "math"

/**
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete at most two transactions.
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

还是股票问题，但是最多只能买卖2次

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
*/

/**
0 <= i < len(prices), 0 <= k <= 2 （交易的次数）, z = 0 | 1（未持有股票|持有股票）
假设认为买了股票就算完成了一次交易
dp[i][1][1] = max(dp[i-1][1][1], - prices[i])
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])

i只和i-1有关系，可以进行简化
dp11 = max(dp11, -prices[i])
dp10 = max(dp10, dp11 + prices[i])
dp21 = max(dp21, dp10 - prices[i])
dp20 = max(dp20, dp21 + prices[i])
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp11, dp10, dp21, dp20 := math.MinInt32, 0, math.MinInt32, 0
	for i := 0; i < len(prices); i++ {
		dp11 = max(dp11, -prices[i])
		dp10 = max(dp10, dp11+prices[i])
		dp21 = max(dp21, dp10-prices[i])
		dp20 = max(dp20, dp21+prices[i])
	}
	return dp20
}
