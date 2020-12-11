package leetcode

import "fmt"

/**
罗马数字转整数
https://leetcode-cn.com/problems/roman-to-integer/
*/
func main() {
	fmt.Print(romanToInt("III"))
}

func romanToInt(s string) int {
	str := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	numMap := make(map[string]int)
	for i := 0; i < len(str); i++ {
		numMap[str[i]] = nums[i]
	}
	res := 0
	for i := 0; i < len(s); {
		if i != len(s)-1 {
			if val, ok := numMap[s[i:i+2]]; ok {
				res += val
				i += 2
				continue
			}
		}
		val := numMap[string(s[i])]
		res += val
		i++
	}
	return res
}
