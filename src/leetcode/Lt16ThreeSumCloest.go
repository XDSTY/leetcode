package leetcode

import (
	"fmt"
	"math"
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
*/

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 1}, 4))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	minCloest := math.MaxInt32
	len := len(nums)
	for i := 0; i < len-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		x := len - 1
		for j < x {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			absVal := abs(target - (nums[i] + nums[j] + nums[x]))
			if absVal == 0 {
				return nums[i] + nums[j] + nums[x]
			}
			if absVal < minCloest {
				minCloest = absVal
				res = nums[i] + nums[j] + nums[x]
			}
			if nums[i]+nums[j]+nums[x] < target {
				j++
			} else {
				x--
			}
		}
	}
	return res
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
