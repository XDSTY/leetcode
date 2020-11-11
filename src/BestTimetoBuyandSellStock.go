package main

import (
	"math"
)

/**
Say you have an array for which the ith element is the price of a given stock on day i.
给你一个数组，数组的每一个元素表示元素下标当前的股票价格
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),
design an algorithm to find the maximum profit.
如果你只能进行一次股票交易，设计一个算法找出最大收益
Note that you cannot sell a stock before you buy one.
不能再股票买入之前售出
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
思路：动态的寻找最小值，与后面比它大的进行比较
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
*/
func maxProfitSign(prices []int) int {
	max, minEle := 0, math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < minEle {
			minEle = prices[i]
		} else {
			if prices[i]-minEle > max {
				max = prices[i] - minEle
			}
		}
	}
	return max
}

/**
dp 的基本形式
# 初始化 base case
dp[0][0][...] = base
# 进行状态转移
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 求最值(选择1，选择2...)
对于股票来说：
状态1: i天数；状态2: k交易次数；状态3: 股票状态0没有持有，1持有
选择：买入、卖出、不买入或继续持有股票reset
对于本题来说，状态转移方程为
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + price[i])
dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - price[i])
dp[i-1][0][0] 是等于0的，因为只能进行一次交易，所以进行交易之前一直未0
因为只能进行一次交易，所以k为1，状态转移方程可以进行简化
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price[i])
dp[i][1] = max(dp[i-1][1], -price[i])
*/
func maxProfitSignWithDp(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(dp); i++ {
		// 第一天只能进行买入，不能卖出
		if i == 0 {
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(dp)-1][0]
}

func max(val1, val2 int) int {
	if val1 < val2 {
		return val2
	}
	return val1
}

/**
还可以对上面的代码进行优化，i 只 用到了 i-1的值，可以去掉dp数组，使用两个变量代替
*/
func maxProfitSignWithDpOne(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dpI0 := 0
	dpI1 := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dpI1 = -prices[i]
			continue
		}
		dpI0 = max(dpI0, dpI1+prices[i])
		dpI1 = max(dpI1, -prices[i])
	}
	return dpI0
}
