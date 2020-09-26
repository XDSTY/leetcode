package main

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
*/
func main() {
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	arr := make([]int, len(s))
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			arr[idx] = int(s[i])
			idx++
		} else {
			if idx == 0 || !pattern(arr[idx-1], int(s[i])) {
				return false
			}
			idx--
		}
	}
	return idx == 0
}

func pattern(c1, c2 int) bool {
	if c1 == '(' {
		return c2 == ')'
	}
	if c1 == '[' {
		return c2 == ']'
	}
	if c1 == '{' {
		return c2 == '}'
	}
	return false
}
