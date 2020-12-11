package leetcode

import "fmt"

/**
Given a collection of distinct integers, return all possible permutations.
Example:
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
思路：回溯法
*/
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	perm(nums, 0, &res)
	return res
}

func perm(nums []int, idx int, res *[][]int) {
	if idx == len(nums) {
		tArr := make([]int, len(nums))
		copy(tArr, nums)
		*res = append(*res, tArr)
		return
	}
	//nums[0:idx-1]是已经确定的值，现在要确定第idx的值，第idx的值可以使所有未使用过的值
	//依次的尝试使用没有用过的值
	for i := idx; i < len(nums); i++ {
		nums[idx], nums[i] = nums[i], nums[idx]
		perm(nums, idx+1, res)
		nums[idx], nums[i] = nums[i], nums[idx]
	}
}
