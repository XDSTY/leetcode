package leetcode

import "fmt"

/**
Given the root of a binary tree, return the length of the longest path,
where each node in the path has the same value. This path may or may not pass through the root.
给一个二叉树，返回具有相同值的路径的长度，路径可能不会经过根节点
The length of the path between two nodes is represented by the number of edges between them.
返回的是路径里面的边的数量
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-univalue-path
*/

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	var findPath func(root *TreeNode) int
	findPath = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		tmpLen := 0
		leftLen := findPath(root.Left)
		rightLen := findPath(root.Right)
		if root.Left != nil && root.Left.Val == root.Val {
			tmpLen = leftLen + 1
		}
		if root.Right != nil {
			if root.Right.Val == root.Val {
				if tmpLen+1+rightLen > maxLen {
					maxLen = tmpLen + 1 + rightLen
				}
				if rightLen+1 >= tmpLen {
					tmpLen = rightLen + 1
				}
			}
		}
		if tmpLen > maxLen {
			maxLen = tmpLen
		}
		return tmpLen
	}
	findPath(root)
	return maxLen
}

func main() {
	node3 := TreeNode{Val: 5}
	node2 := TreeNode{Val: 5, Right: &node3}
	node1 := TreeNode{Val: 5, Right: &node2}
	fmt.Println(longestUnivaluePath(&node1))
}
