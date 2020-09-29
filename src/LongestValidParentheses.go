package main

import "fmt"

/**
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
给定一个只含有'('和')'的字符串，找到最长的可以匹配的子串
Example 1:
Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:
Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses

一看就是dp 艹
只有碰到')'时才需要进行判断
1)、s[i]为')'且s[i-1]为'('，则dp[i] = dp[i - 2] + 2
2)、s[i]为')'且s[i-1]也为')'，要为s[i]找到匹配的'('，s[i]匹配的'('应该在s[i - dp[i-1] - 1]，所以只要s[i - dp[i-1] - 1]为'('，s[i]就能匹配
	dp[i] = dp[i - 1] + dp[i - dp[i - 1] - 2] + 2

这题还是用栈更加简单
栈中存放字符串字符的下标，当碰到'('直接入栈，碰到')'则找栈顶的值是否匹配，匹配则将栈顶出栈，否则继续入栈
*/

func main() {
	fmt.Println(longestValidParenthesesWithDp(")()())"))
}

func longestValidParenthesesWithDp(s string) int {
	dp := make([]int, len(s)+1)
	maxLen := 0
	for i := 2; i <= len(s); i++ {
		if s[i-1] == ')' && s[i-2] == '(' {
			dp[i] = dp[i-2] + 2
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		} else if s[i-1] == ')' && s[i-2] == ')' && i-dp[i-1]-2 >= 0 && s[i-dp[i-1]-2] == '(' {
			dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}

func longestValidParenthesesWithStack(s string) int {
	stack := make([]int, len(s)+1)
	maxLen := 0
	idx := 0
	for i := range s {
		if s[i] == '(' {
			stack[idx] = i
			idx++
		} else {
			if idx > 0 && s[stack[idx-1]] == '(' {
				idx--
				if idx == 0 {
					maxLen = i + 1
				} else if i-stack[idx-1]+1 > maxLen {
					maxLen = i - stack[idx-1]
				}
			} else {
				stack[idx] = i
				idx++
			}
		}
	}
	return maxLen
}
