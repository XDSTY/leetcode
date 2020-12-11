package leetcode

import "fmt"

/**
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
*/
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	row := len(matrix)
	column := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < column-i; j++ {
			matrix[i][j], matrix[column-j-1][row-i-1] = matrix[column-j-1][row-i-1], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
	}
}
