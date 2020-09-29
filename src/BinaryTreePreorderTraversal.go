package main

import "collections"

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preOrderTraversalWithRecur(root, &res)
	return res
}

/**
非递归前序遍历
*/
func preOrderTraversalWithRecur(root *TreeNode, arr *[]int) {
	stack := collections.NewStack()
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			*arr = append(*arr, node.Val)
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop().(*TreeNode)
		node = node.Right
	}
}
