package main

/**
Given the root of a binary tree, return the sum of every tree node's tilt.
给一个二叉树，返回二叉树的坡度
The tilt of a tree node is the absolute difference between the sum of all left subtree node values
and all right subtree node values. If a node does not have a left child,
then the sum of the left subtree node values is treated as 0. The rule is similar if there the node does not have a right child.
坡度是二叉树的每个节点的左右子树的差值之和
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-tilt
*/
func findTilt(root *TreeNode) int {
	res := 0
	var findNodeSum func(node *TreeNode, res *int) int
	findNodeSum = func(node *TreeNode, res *int) int {
		if node == nil {
			return 0
		}
		left := findNodeSum(node.Left, res)
		right := findNodeSum(node.Right, res)
		sub := left - right
		if sub < 0 {
			sub = -sub
		}
		*res += sub
		return node.Val + left + right
	}
	findNodeSum(root, &res)
	return res
}
