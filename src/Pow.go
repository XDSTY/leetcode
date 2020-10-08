package main

import "fmt"

/**
Implement pow(x, n), which calculates x raised to the power n (i.e. xn).
实现pow(x,n)
Example 1:
Input: x = 2.00000, n = 10
Output: 1024.00000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/powx-n
思路：从低到高判断n的二进制位是否为1
*/

func main() {
	fmt.Println(myPow(2, 10))
}

func myPow(x float64, n int) float64 {
	if n > 0 {
		return multi(x, n)
	}
	return 1 / multi(x, -n)
}

func multi(x float64, n int) float64 {
	res := 1.0
	tmp := x
	for n > 0 {
		// 判断当前位是否为1
		if n%2 == 1 {
			res *= tmp
		}
		tmp *= tmp
		n >>= 1
	}
	return res
}
