package main

import "math"

/**
Given a Binary Search Tree (BST) with the root node root,
return the minimum difference between the values of any two different nodes in the tree.
给一个二叉搜索树，找到任意两点之间最小的差值

思路：因为是二叉搜索树，所以只需要判断相邻节点的差值就可以
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
*/

func minDiffInBST(root *TreeNode) int {
	minVal := math.MaxInt32
	var preNode *TreeNode

	var getMin func(node *TreeNode)
	getMin = func(node *TreeNode) {
		if node == nil {
			return
		}
		getMin(node.Left)
		if preNode == nil {
			preNode = node
		} else {
			if node.Val-preNode.Val < minVal {
				minVal = node.Val - preNode.Val
			}
			preNode = node
		}
		getMin(node.Right)
	}
	getMin(root)
	return minVal
}

func inOrderWithStack() {

}
