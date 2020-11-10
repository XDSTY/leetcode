package main

import "fmt"

/**
Given an integer array nums, return the number of range sums that lie in [lower, upper] inclusive.
Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j (i ≤ j), inclusive.
给一个整数数组，返回区间和在[lower,upper]之间的数量
区间的和S(i, j)表示数组[i,j]之间的和

Input: nums = [-2,5,-1], lower = -2, upper = 2,
Output: 3
Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective sums are: -2, -1, 2.

思路：新建数组sums，sum[i]表示[0,i]之间的和，问题变成sums[j] - sums[i]的值在[lower, upper]之间的数量
	  再使用归并算法对sums数组进行排序，排序时[left, mid, right]当sums[j] (j >= mid && j =< right) 和sums[i] (i>=left && i <= mid) 满足
      sums[j] - sums[i]大于lower并小于upper，则[i,j]就是符合的区间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-of-range-sum
*/
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = nums[i] + sums[i]
	}

	var mergeSort func(left, right int) int
	mergeSort = func(left, right int) int {
		if left == right {
			return 0
		}
		mid := (left + right) / 2
		res := mergeSort(left, mid) + mergeSort(mid+1, right)

		// 统计符合的数量
		l, r, i := mid+1, mid+1, left
		for i <= mid {
			// 在[mid,right]里面寻找
			// 找到=>lower的左区间下标
			for l <= right && sums[l]-sums[i] < lower {
				l++
			}
			// 找到<=upper的右区间下标
			for r <= right && sums[r]-sums[i] <= upper {
				r++
			}
			res += r - l
			i++
		}

		// 合并数组
		tmpArr := make([]int, right-left+1)
		li, ri, i := left, mid+1, 0
		for ; li <= mid || ri <= right; i++ {
			if li <= mid && ri <= right {
				if sums[li] < sums[ri] {
					tmpArr[i] = sums[li]
					li++
				} else {
					tmpArr[i] = sums[ri]
					ri++
				}
			} else if li <= mid {
				tmpArr[i] = sums[li]
				li++
			} else {
				tmpArr[i] = sums[ri]
				ri++
			}
		}
		for i, j := left, 0; i <= right; i++ {
			sums[i] = tmpArr[j]
			j++
		}
		return res
	}
	return mergeSort(0, len(sums)-1)
}

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
