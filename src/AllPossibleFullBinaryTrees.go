package main

/**
A full binary tree is a binary tree where each node has exactly 0 or 2 children.
一个满二叉树是每个节点有0个或者2个子节点的二叉树
Return a list of all possible full binary trees with N nodes.  Each element of the answer is the root node of one possible tree.
返回具有n个节点的所有满二叉树
Each node of each tree in the answer must have node.val = 0.
每个节点的值都是0
You may return the final list of trees in any order.

思路：如果当前节点为根节点的树是满二叉树，则该节点的左右节点都是满二叉树
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-possible-full-binary-trees
*/

func allPossibleFBT(N int) []*TreeNode {
	if N == 0 {
		return nil
	}
	if N == 1 {
		return []*TreeNode{new(TreeNode)}
	}
	// 注意起点是1，满二叉树左边或右边必须有一个节点
	roots := make([]*TreeNode, 0)
	for i := 1; i <= N-2; i += 2 {
		l := allPossibleFBT(i)
		r := allPossibleFBT(N - 1 - i)
		for _, left := range l {
			for _, right := range r {
				root := &TreeNode{Left: left, Right: right}
				roots = append(roots, root)
			}
		}
	}
	return roots
}
