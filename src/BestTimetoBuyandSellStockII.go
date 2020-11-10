package main

/**
Say you have an array prices for which the ith element is the price of a given stock on day i.
给你一个数组，数组中的，每个元素的下标代表着天数，值代表着这天的股票价格
Design an algorithm to find the maximum profit.
You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
设计一个算法找到最大的收益，可以进行多次的交易
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
不能同时进行交易

思路：找升序段？
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
