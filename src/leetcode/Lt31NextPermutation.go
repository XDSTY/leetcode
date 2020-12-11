package leetcode

import "fmt"

/**
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
The replacement must be in place and use only constant extra memory.
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
Example 1:
Input: nums = [1,2,3]
Output: [1,3,2]
Example 2:
Input: nums = [3,2,1]
Output: [1,2,3]
Example 3:
Input: nums = [1,1,5]
Output: [1,5,1]
Example 4:
Input: nums = [1]
Output: [1]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation

下一个排列
解题思路：https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
	1、从后向前查找第一个相邻升序的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
	2、在 [j,end) 从后向前查找第一个满足 A[i] < A[k] 的 k。A[i]、A[k] 分别就是上文所说的「小数」、「大数」
	3、将 A[i] 与 A[k] 交换
	4、可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
	5、如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 4
*/
func main() {
	nums := make([]int, 3)
	nums[0] = 5
	nums[1] = 1
	nums[2] = 1
	//nums[3] = 4
	//nums[4] = 6
	//nums[5] = 5
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	j := len(nums) - 1
	for ; j > 0 && nums[j-1] >= nums[j]; j-- {
	}
	if j > 0 {
		i := j - 1
		k := len(nums) - 1
		for ; k >= j && nums[k] <= nums[i]; k-- {
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	for x := len(nums) - 1; x > j; x-- {
		nums[j], nums[x] = nums[x], nums[j]
		j++
	}
}
