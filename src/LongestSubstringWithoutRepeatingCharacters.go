package main

import (
	"fmt"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/
func main() {
	fmt.Println(int('z'))
	fmt.Print(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	chars := make([]int, 256)
	maxLen := 0
	tmpLen := 0
	tmpI := 0
	tmpJ := 0
	for i := 0; i < len(s); i++ {
		tmpJ = i
		if chars[s[i]] == 0 {
			tmpLen++
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
			chars[s[i]] = 1
		} else {
			for ti := tmpI; ti <= tmpJ; ti++ {
				if s[ti] != s[i] {
					chars[s[ti]] = 0
				} else if s[ti] == s[i] {
					tmpI = ti + 1
					tmpLen = tmpJ - tmpI + 1
					break
				}
			}
		}
	}
	return maxLen
}
