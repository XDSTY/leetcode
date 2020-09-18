package main

import "fmt"

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
输入: 121
输出: true
*/
func main() {
	fmt.Print(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res := 0
	tmp := x
	for tmp > 0 {
		res *= 10
		res += tmp % 10
		tmp /= 10
	}
	return x == res
}
