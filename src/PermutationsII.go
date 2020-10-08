package main

import (
	"fmt"
	"sort"
)

/**
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
含有重复数字的全排列
思路：回溯法加上一个标记已经访问过的数组
*/
func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1}))
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	vis := make([]bool, len(nums))
	tarr := make([]int, 0)
	var backtrack func(int)
	sort.Ints(nums)
	backtrack = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int(nil), tarr...))
			return
		}
		for i := 0; i < len(nums); i++ {
			// 当前数字已被访问 或者 当前的值和前一个值相等并且前一个值没有被访问（因为虽然这次回溯前一个值没有被访问吗，但是其他的回溯里面会有先访问前一个值再访问当前值得情况）
			// 所以当前一个相同的值被访问后，当前值才能被访问
			if vis[i] || (i > 0 && nums[i] == nums[i-1] && !vis[i-1]) {
				continue
			}
			tarr = append(tarr, nums[i])
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			tarr = tarr[0 : len(tarr)-1]
		}
	}
	backtrack(0)
	return res
}
