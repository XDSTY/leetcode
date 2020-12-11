package leetcode

import (
	"fmt"
	"strconv"
)

/**
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
给两个string类型的整数，获取他们的乘机

Example 1:
Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:
Input: num1 = "123", num2 = "456"
Output: "56088"

123
456
----
3*6 + 20*6 + 100*6 + 50 * 3 + 50 * 20 + 50 * 100 + 400*3 + 400*20+ 400*100 =
18,120,600,150,1000,5000,1200,8000,40000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
*/
func main() {
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1 := len(num1)
	len2 := len(num2)
	sumArr := make([]int, len1+len2)
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			tsum := (num1[i] - '0') * (num2[j] - '0')
			sumArr[len1-i+len2-j-2] += int(tsum)
		}
	}
	tmp := 0
	res := ""
	for i := 0; i < len(sumArr); i++ {
		sumArr[i] += tmp
		tmp = sumArr[i] / 10
		sumArr[i] %= 10
	}
	if sumArr[len(sumArr)-1] > 0 {
		res += strconv.Itoa(sumArr[len(sumArr)-1])
	}
	for i := len(sumArr) - 2; i >= 0; i-- {
		res += strconv.Itoa(sumArr[i])
	}
	return res
}
