package leetcode

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	nodeArr := make([]*TreeNode, 0)
	nodeArr = append(nodeArr, root)
	for len(nodeArr) > 0 {
		arr := make([]int, 0)
		newNodeArr := make([]*TreeNode, 0)
		for i := 0; i < len(nodeArr); i++ {
			arr = append(arr, nodeArr[i].Val)
			if nodeArr[i].Left != nil {
				newNodeArr = append(newNodeArr, nodeArr[i].Left)
			}
			if nodeArr[i].Right != nil {
				newNodeArr = append(newNodeArr, nodeArr[i].Right)
			}
		}
		res = append(res, arr)
		nodeArr = newNodeArr
	}
	return res
}
