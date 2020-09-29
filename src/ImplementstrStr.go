package main

import "fmt"

/**
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
*/
func main() {
	//"mississippi"
	//"a"
	fmt.Println(strStr("aaa", "aaa"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	haystackIdx := 0
	for haystackIdx < len(haystack) {
		for ; haystackIdx < len(haystack) && haystack[haystackIdx] != needle[0]; haystackIdx++ {
		}
		idx := 0
		for tidx := haystackIdx; tidx < len(haystack) && idx < len(needle); idx++ {
			if haystack[tidx] != needle[idx] {
				break
			}
			tidx++
		}
		if idx == len(needle) {
			return haystackIdx
		}
		haystackIdx++
	}
	return -1
}
