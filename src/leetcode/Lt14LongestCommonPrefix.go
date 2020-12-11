package leetcode

import "fmt"

/**
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
https://leetcode-cn.com/problems/longest-common-prefix/
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
