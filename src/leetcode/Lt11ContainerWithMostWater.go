package leetcode

import (
	"fmt"
)

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
输入：[1,8,6,2,5,4,8,3,7]
输出：49

思路：双指针，计算后移动值较小的指针
*/

func main() {
	fmt.Print(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		if (j-i)*min(height[i], height[j]) > max {
			max = (j - i) * min(height[i], height[j])
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}

func min(v, v1 int) int {
	if v < v1 {
		return v
	}
	return v1
}
