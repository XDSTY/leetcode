package leetcode

/**
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
给定n，返回有多少种二叉搜索树
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
/**
树的思想一般是将问题分解为左右子树的问题
该问题可以固定根节点，左节点组成的二叉搜索树*右节点组成的二叉搜索树数量，所以可以使用dp的方式
*/
func numTrees(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}
