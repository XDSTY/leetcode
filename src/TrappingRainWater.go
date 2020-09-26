package main

import "fmt"

/**
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.
一组非负数的数字代表一个梯度图，每个柱子的宽度是1，计算这个梯度图可以乘下多少水

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
*/

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	stack := make([]int, len(height))
	idx := 0
	res := 0
	for i := 0; i < len(height); i++ {
		for idx > 0 && height[stack[idx-1]] < height[i] {
			val := stack[idx-1]
			idx--
			if idx == 0 {
				break
			}
			distance := i - stack[idx-1] - 1
			hei := getMin(height[i], height[stack[idx-1]]) - height[val]
			res += hei * distance
		}
		stack[idx] = i
		idx++
	}
	return res
}

func getMin(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
