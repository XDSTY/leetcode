package main

/**
The Tribonacci sequence Tn is defined as follows:

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-th-tribonacci-number
*/

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	val1, val2, val3, val4 := 0, 1, 1, 0
	for i := 3; i <= n; i++ {
		val4 = val1 + val2 + val3
		val1 = val2
		val2 = val3
		val3 = val4
	}
	return val3
}
