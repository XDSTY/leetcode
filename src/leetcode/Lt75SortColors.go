package leetcode

import (
	"fmt"
)

/**
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent,
with the colors in the order red, white, and blue.
Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
给一个包含红、白、蓝对象的数组，将他们排序成相邻的数组，不能使用额外的空间

Follow up:
Could you solve this problem without using the library's sort function?
Could you come up with a one-pass algorithm using only O(1) constant space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
思路：可以使用qsort的partition来做，找到0的最右边的值，和2的最左边的值，则数组就变成了 0 1 2的形式
双指针：p0指针指向0的下标，p2指针指向2的下标
*/

func main() {
	nums := []int{2, 0, 1}
	sortColorsWithPoint(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i, v := range nums {
		if v == 0 {
			nums[0], nums[i] = nums[i], nums[0]
		}
	}

	var partition func(left, right int) int
	partition = func(left, right int) int {
		flag := nums[left]
		for left < right {
			for left < right && nums[right] > flag {
				right--
			}
			for left < right && nums[left] <= flag {
				left++
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
		return left
	}

	left := 0
	if nums[0] == 0 {
		left = partition(0, len(nums)-1) + 1
	}
	for j := len(nums) - 1; left < j; {
		for left < j && nums[j] == 2 {
			j--
		}
		for left < j && nums[left] == 1 {
			left++
		}
		if left < j {
			nums[j], nums[left] = nums[left], nums[j]
		}
	}
}

func sortColorsWithPoint(nums []int) {
	if len(nums) < 2 {
		return
	}
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i < p2 && nums[i] == 2 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		}
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		}
	}
}
