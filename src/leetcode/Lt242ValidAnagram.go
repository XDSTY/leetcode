package leetcode

import "fmt"

/**
Given two strings s and t , write a function to determine if t is an anagram of s.
判断t是不是s的字母异位词
Example 1:
Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := [26]int{}
	for _, v := range s {
		chars[v-'a'] += 1
	}
	tchars := [26]int{}
	for _, v := range t {
		tchars[v-'a'] += 1
	}
	for i := 0; i < len(chars); i++ {
		if chars[i] != tchars[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("aac", "csa"))
}
