package main

import "fmt"

/**
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
*/

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left, right, top, down := 0, n-1, 0, n-1
	num := 1
	for true {
		for j := left; j <= right; j++ {
			matrix[top][j] = num
			num++
		}
		top++
		if top > down {
			break
		}

		for i := top; i <= down; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		if right < left {
			break
		}

		for j := right; j >= left; j-- {
			matrix[down][j] = num
			num++
		}
		down--
		if down < top {
			break
		}

		for i := down; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}
	return matrix
}

func main() {
	matrix := generateMatrix(1)
	fmt.Println(matrix)
}
