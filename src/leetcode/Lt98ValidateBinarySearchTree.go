package leetcode

/**
Given a binary tree, determine if it is a valid binary search tree (BST).
给一个二叉树，【判断是不是二叉搜索树
Assume a BST is defined as follows:
二叉搜索树的定义如下
The left subtree of a node contains only nodes with keys less than the node's key.
左节点的值小于当前节点
The right subtree of a node contains only nodes with keys greater than the node's key.
右节点的值大于当前节点
Both the left and right subtrees must also be binary search trees.


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValidBST(root *TreeNode) bool {
	preNode := &TreeNode{}
	preNode = nil
	var validBST func(node *TreeNode) bool
	validBST = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		res := validBST(node.Left)
		if !res || (preNode != nil && preNode.Val >= node.Val) {
			return false
		}
		preNode = node
		res = validBST(node.Right)
		return res
	}
	return validBST(root)
}
