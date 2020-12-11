package leetcode

import "fmt"

/**
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
Note: A leaf is a node with no children.
给一个二叉树和一个值，找到所有二叉树的根节点到叶子节点的和等于给定值的路劲的集合
Example:
Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]
*/

func main() {
	leftNode := TreeNode{Val: 2}
	rightNode := TreeNode{Val: -3}
	root := TreeNode{Val: -2, Left: &leftNode, Right: &rightNode}
	fmt.Println(pathSum(&root, -5))
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	pathSumF(root, sum, 0, &res, []int{})
	return res
}

func pathSumF(node *TreeNode, sum, tmp int, res *[][]int, arr []int) {
	if node == nil {
		return
	}
	if node.Right == nil && node.Left == nil && tmp+node.Val == sum {
		*res = append(*res, append(arr, node.Val))
		return
	}
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)
	tmpArr = append(tmpArr, node.Val)
	pathSumF(node.Left, sum, tmp+node.Val, res, tmpArr)
	pathSumF(node.Right, sum, tmp+node.Val, res, tmpArr)
}
