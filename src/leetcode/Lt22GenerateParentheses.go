package leetcode

import "fmt"

/**

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

https://leetcode-cn.com/problems/generate-parentheses/
就是n个数有多少种出栈方式，对于第n个数字，有出栈和不出栈两种选择
*/

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs(n, 0, 0, "", &res)
	return res
}

func dfs(n, open, close int, re string, res *[]string) {
	if open == n && close == n {
		*res = append(*res, re)
		return
	}
	if open < n {
		dfs(n, open+1, close, re+"(", res)
	}
	if close < open {
		dfs(n, open, close+1, re+")", res)
	}
}
