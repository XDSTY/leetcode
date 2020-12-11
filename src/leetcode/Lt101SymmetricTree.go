package leetcode

/**
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var symmetric func(left, right *TreeNode) bool
	symmetric = func(left, right *TreeNode) bool {
		if left != nil && right != nil {
			return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
		}
		if left == nil && right == nil {
			return true
		}
		return false
	}
	return symmetric(root.Left, root.Right)
}
