package leetcode

import (
	"fmt"
	"sort"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
思路：3sum转化为2sum，固定一个值，在剩余的值里面进行查找
*/
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// 一种思路，先排序再将问题转化为2sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	len := len(nums)
	for i := 0; i < len-2; i++ {
		// 如果当前值和上一个值相同，只会产生重复的结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		x := len - 1
		for j < x {
			//如果当前值和上一个值相同，只会产生重复的结果
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[j]+nums[x] == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[x]})
				j++
				x--
			} else if nums[j]+nums[x] < -nums[i] {
				j++
			} else {
				x--
			}
		}
	}
	return res
}
