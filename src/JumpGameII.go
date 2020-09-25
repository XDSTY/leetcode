package main

import "fmt"

/**
Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
找到可以到达最后位置的最少跳跃次数
贪心算法：每次都找到跳到最远的边界，贪心

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-ii
*/
func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	maxSize := nums[0]
	step := 0
	position := 0
	for position < len(nums) {
		if maxSize >= len(nums)-1 {
			return step + 1
		}
		tSize := maxSize
		for j := position + 1; j <= tSize; j++ {
			if nums[j]+j > maxSize {
				maxSize = nums[j] + j
				position = j
			}
		}
		step++
	}
	return step
}
