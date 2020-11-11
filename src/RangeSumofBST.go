package main

/**
Given the root node of a binary search tree, return the sum of values of all nodes with a value in the range [low, high].
给一个二叉搜索树，返回和在[low,high]之间的和

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-of-bst
*/

func rangeSumBST(root *TreeNode, low int, high int) int {
	var sumBST func(node *TreeNode) int
	findFlag := false
	sumBST = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := sumBST(node.Left)
		if findFlag && (node.Val < low || node.Val > high) {
			return sum
		}
		if node.Val >= low && node.Val <= high {
			sum += node.Val
			findFlag = true
		}
		sum += sumBST(node.Right)
		return sum
	}
	return sumBST(root)
}
