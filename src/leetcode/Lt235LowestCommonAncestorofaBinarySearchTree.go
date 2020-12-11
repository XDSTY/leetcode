package leetcode

/**
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.
给定一个二叉搜索树和两个节点，找到两个节点的最近的公共父节点
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if q.Val > root.Val && p.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
