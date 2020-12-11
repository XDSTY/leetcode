package leetcode

/**
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q != nil && p != nil {
		if q.Val != p.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	if q == nil && p == nil {
		return true
	}
	return false
}
