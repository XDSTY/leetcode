package main

/**
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
注意是叶子节点

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
*/

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return leaves(root.Left, 1) + leaves(root.Right, 2)
}

func leaves(node *TreeNode, typeNode int) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil && typeNode == 1 {
		return node.Val
	}
	return leaves(node.Left, 1) + leaves(node.Right, 2)
}
