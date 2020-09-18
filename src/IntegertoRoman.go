package main

import "fmt"

/**
整数转罗马数字
https://leetcode-cn.com/problems/integer-to-roman/
*/
func main() {
	fmt.Print(intToRoman(1994))
}

func intToRoman(num int) string {
	str := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	res := ""
	for i := len(nums) - 1; i >= 0; i-- {
		for num-nums[i] >= 0 {
			res += str[i]
			num -= nums[i]
		}
	}
	return res
}
