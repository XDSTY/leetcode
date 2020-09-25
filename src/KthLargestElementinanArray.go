package main

import "fmt"

/**
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order,
not the kth distinct element.
在数组中找到第k大的数

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4

来源：力扣（LeetCode）
*/

func main() {
	fmt.Print(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	// 构建最大堆
	for i := len(nums) / 2; i >= 0; i-- {
		buildMaxHeap(nums, i, len(nums))
	}
	for i := 0; i < k; i++ {
		swap(nums, len(nums)-1-i, 0)
		buildMaxHeap(nums, 0, len(nums)-i-1)
	}
	return nums[len(nums)-k]
}

func buildMaxHeap(nums []int, idx int, len int) {
	maxIdx := idx
	if 2*idx+1 < len && nums[maxIdx] < nums[2*idx+1] {
		maxIdx = 2*idx + 1
	}
	if 2*idx+2 < len && nums[maxIdx] < nums[2*idx+2] {
		maxIdx = 2*idx + 2
	}
	if idx != maxIdx {
		swap(nums, idx, maxIdx)
		buildMaxHeap(nums, maxIdx, len)
	}
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
