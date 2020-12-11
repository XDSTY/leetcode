package leetcode

import "fmt"

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
https://leetcode-cn.com/problems/longest-palindromic-substring/
思路：s[i+1 ... j-1]是回文且s[i]==s[j]   或者 j - i < 3且s[j] == s[i]
*/
func main() {
	fmt.Print(longestPalindrome("ac"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := s[0:1]
	for r := 0; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if s[r] == s[l] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
				if r-l+1 > len(res) {
					res = s[l : r+1]
				}
			}
		}
	}
	return res
}
