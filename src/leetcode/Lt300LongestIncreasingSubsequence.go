package leetcode

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence

思路：dp[i] 表示以nums[i]为结尾的最大升序的长度
	dp[i] = dp[j] + 1 ,  0 <= j <i，nums[i] > nums[j]
*/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return dp[len(nums)-1]
}

func lengthOfLIS1(nums []int) int {
	// d[i] 表示长度为i时，子序列的末尾的值
	// num[j] 如果大于d[..]的末尾的值，则d[len++]=num[j]，如果num[j]小于d[..]末尾的值，则使用二分法找到自己在d[..]的位置
	d := make([]int, len(nums)+1)
	length := 1
	d[length] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[length] {
			length++
			d[length] = nums[i]
		} else {
			if nums[i] < d[1] {
				d[1] = nums[i]
			} else {
				left, right := 1, length
				for left <= right {
					mid := left + (right-left)/2
					if d[mid] < nums[i] {
						left = mid + 1
					} else {
						right = mid - 1
					}
				}
				d[left] = nums[i]
			}
		}
	}
	return length
}
