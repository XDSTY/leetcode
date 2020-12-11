package leetcode

import "fmt"

/**
Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Determine if you are able to reach the last index.
给定一个非负整数数组，每个元素是你能往后跳的最大长度，现在你在数组的第一个位置，判断是否可以到达最后一个位置

Input: nums = [2,3,1,1,4]
Output: true
Input: nums = [3,2,1,0,4]
Output: false
思路：贪心，维护一个前n项数字可以到达的最大位置，只要最大位置>length，就返回true
*/

func main() {
	fmt.Print(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	maxSize := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxSize >= len(nums)-1 {
			return true
		}
		if i <= maxSize && maxSize < i+nums[i] {
			maxSize = i + nums[i]
		}
	}
	return maxSize >= len(nums)-1
}
