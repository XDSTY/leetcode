package main

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
	fmt.Print(longestPalindrome("aaaa"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[i-1][j+1] == true) {
				dp[i][j] = true
				if i-j+1 > len(res) {
					res = s[j : i+1]
				}
			}
		}
	}
	return res
}
