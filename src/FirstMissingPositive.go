package main

import "fmt"

/**
Given an unsorted integer array, find the smallest missing positive integer.
Example 1:
Input: [1,2,0]
Output: 3
Example 2:
Input: [3,4,-1,1]
Output: 2
Example 3:
Input: [7,8,9,11,12]
Output: 1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
思路：1、使用map存储每个数字，从1开始查询map，如果map中没有当前的数，则当前数就是结果
2、如果数组中的值>1&&<N(数组长度)，则将该数放入数组的对应位置，如果数组的当前值不等于下标+1，则当前下标+1就是结果
*/
func main() {
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
