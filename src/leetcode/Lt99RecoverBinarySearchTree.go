package leetcode

import (
	"collections"
)

/**
You are given the root of a binary search tree (BST), where exactly two nodes of the tree were swapped by mistake.
Recover the tree without changing its structure.
给一个二叉搜索树，其中两个节点错误的交换了，在不改变结构的情况下，恢复这棵树
Follow up: A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recover-binary-search-tree

思路：比如说一颗二叉搜索树的中序遍历为 1 2 3 4 5
交换1和4则变成，4 2 3 1 5，则只需要中序遍历找到4和1两个节点即可，可以看到4到2下降了一次，3到1又下降了一次，可以记录下来
*/
func recoverTree(root *TreeNode) {
	var x *TreeNode
	var y *TreeNode
	var preNode *TreeNode

	//非递归中序遍历
	stack := collections.NewStack()
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop().(*TreeNode)
		if preNode != nil && node.Val < preNode.Val {
			y = node
			if x == nil {
				x = preNode
			} else {
				break
			}
		}
		preNode = node
		node = node.Right
	}
	x.Val, y.Val = y.Val, x.Val
}
