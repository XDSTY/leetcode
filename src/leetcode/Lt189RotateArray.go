package leetcode

/**
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
https://leetcode-cn.com/problems/rotate-array/
*/
/**
解法1：使用环状数组
*/
func rotateArray(nums []int, k int) {
	length := len(nums)
	k = k % length
	count := 0
	for start := 0; start < length && count < length; start++ {
		tmp := nums[start]
		val := nums[start]
		idx := start

		for count < length {
			newIdx := (idx + k) % length
			tmp = nums[newIdx]
			nums[newIdx] = val
			val = tmp
			count++
			if start == (idx+k)%length {
				break
			}
			idx = (idx + k) % length
		}
	}
}

/**
解法二：使用翻转
假设 n=7 且 k=3 。
原始数组                  : 1 2 3 4 5 6 7
反转所有数字后             : 7 6 5 4 3 2 1
反转前 k 个数字后          : 5 6 7 4 3 2 1
反转后 n-k 个数字后        : 5 6 7 1 2 3 4 --> 结果
*/

func rotateArrayWithReverse(nums []int, k int) {
	var reverse func(l, r int)
	reverse = func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	k = k % len(nums)
	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)
}
