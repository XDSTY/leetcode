package main

import "math"

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
