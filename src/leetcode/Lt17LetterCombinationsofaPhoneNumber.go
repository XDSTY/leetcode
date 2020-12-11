package leetcode

import (
	"fmt"
)

/**
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
*/

var chars = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func main() {
	fmt.Println(letterCombinations(""))
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	arr := make([]int, len(digits))
	dfs(len(digits), 0, digits, &res, arr)
	return res
}

func dfs(length int, curr int, digits string, res *[]string, arr []int) {
	if length == curr {
		*res = append(*res, getStr(arr))
		return
	}
	str := chars[(digits[curr]-'0')-2]
	for i := 0; i < len(str); i++ {
		arr[curr] = int(str[i])
		dfs(length, curr+1, digits, res, arr)
	}
}

func getStr(arr []int) string {
	res := ""
	for i := 0; i < len(arr); i++ {
		res += string(arr[i])
	}
	return res
}
