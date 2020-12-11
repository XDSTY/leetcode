package leetcode

import "fmt"

/**
Given a set of distinct integers, nums, return all possible subsets (the power set).
Note: The solution set must not contain duplicate subsets.
一个没有重复数字的整数集合，找出所有的不重复子集
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets

回溯法：考虑当前值和不考虑当前值，如果都没不考虑，则结果是空，都考虑则结果就是全集

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	subsetsRec(nums, 0, make([]int, len(nums)), 0, &res)
	return res
}

func subsetsRec(nums []int, index int, arr []int, idx int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, append([]int(nil), arr[0:idx]...))
		return
	}
	// 不考虑当前值
	subsetsRec(nums, index+1, arr, idx, res)
	arr[idx] = nums[index]
	idx++
	subsetsRec(nums, index+1, arr, idx, res)
}
