package main

import "fmt"

/**
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?
一个机器人位于mxn表格的左上的位置，机器人可以向下或者向右移动，机器人尝试移动到右下的位置，机器人有多少中方式可以达到右下的位置

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
*/
func main() {
	fmt.Println(uniquePaths(3, 3))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
