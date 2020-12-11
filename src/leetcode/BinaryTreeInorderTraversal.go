package leetcode

import "collections"

/**
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
*/

func inorderTraversal(root *TreeNode) []int {
	stack := collections.NewStack()
	node := root
	arr := make([]int, 0)
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop().(*TreeNode)
		arr = append(arr, node.Val)
		node = node.Right
	}
	return arr
}
