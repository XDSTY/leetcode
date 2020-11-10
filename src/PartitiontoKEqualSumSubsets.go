package main

/**
Given an array of integers nums and a positive integer k,
find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.
给一个整数数组和一个正整数数字k，将数组划分为k个具有相同和的数组

Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets
*/

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k > 0 {
		return false
	}
	div := sum / k

	quickSort(nums, 0, len(nums)-1)

	row := len(nums) - 1
	if nums[row] > div {
		return false
	}
	for row >= 0 && nums[row] == div {
		k--
		row--
	}

	groups := make([]int, k)
	var findPartSum func(row int, k int) bool
	findPartSum = func(row int, k int) bool {
		if row < 0 {
			return true
		}
		val := nums[row]
		row--
		for i := 0; i < len(groups); i++ {
			if groups[i]+val <= div {
				groups[i] += val
				if findPartSum(row, k) {
					return true
				}
				groups[i] -= val
			}
		}
		return false
	}
	return findPartSum(row, k)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		partition := partition(nums, left, right)
		quickSort(nums, left, partition)
		quickSort(nums, partition+1, right)
	}
}

func partition(nums []int, left, right int) int {
	flag := nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= flag {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= flag {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = flag
	return i
}
