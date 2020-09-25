package main

import "fmt"

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach,
which is more subtle.
给一组数字，找到一组连续的数字拥有最大的和
如果你已经找到了O(n)的解法，尝试使用更为精妙的分治法求解

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
*/

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	tmpSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if tmpSum < 0 {
			tmpSum = nums[i]
		} else {
			tmpSum += nums[i]
		}
		if tmpSum > maxSum {
			maxSum = tmpSum
		}
	}
	return maxSum
}
