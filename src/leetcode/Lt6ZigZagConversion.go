package leetcode

import (
	"fmt"
	"strings"
)

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
*/
func main() {
	fmt.Print(convert("LEETCODEISHIRING", 4))
}

func convert(s string, numRows int) string {
	if len(s) <= 2 || numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	downLen := 0
	upLen := 0
	for i := 0; i < len(s); i++ {
		if downLen == numRows && upLen == numRows-2 {
			downLen = 0
			upLen = 0
		}
		if downLen < numRows {
			res[downLen] = res[downLen] + string(s[i])
			downLen++
		} else if upLen < numRows-2 {
			res[numRows-2-upLen] = res[numRows-2-upLen] + string(s[i])
			upLen++
		}
	}
	return strings.Join(res, "")
}
