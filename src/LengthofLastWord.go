package main

import "fmt"

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
return the length of last word (last word means the last appearing word if we loop from left to right) in the string.
If the last word does not exist, return 0.
Note: A word is defined as a maximal substring consisting of non-space characters only.
Example:
Input: "Hello World"
Output: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word
*/

func main() {
	fmt.Println(lengthOfLastWord("Hello World     "))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if res == 0 && s[i] == ' ' {
			continue
		}
		if res != 0 && s[i] == ' ' {
			return res
		}
		res++
	}
	return res
}
