package main

import "fmt"

/**
 * 两数之和
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/two-sum
	思路：如果v1+v2 == target，v1在map中，则map[v1 = (target-v2)] != null.
			所以遍历nums，当前val在map中不存在，则map[target-val]=i，val在map中存在，则找到结果
*/
func main() {
	nums := []int{11, 15, 7, 2}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return nil
	}
	var res []int
	var m = make(map[int]int)
	for i, k := range nums {
		if val, exist := m[k]; exist {
			res = append(res, val)
			res = append(res, i)
			return res
		} else {
			m[target-k] = i
		}
	}
	return nil
}
