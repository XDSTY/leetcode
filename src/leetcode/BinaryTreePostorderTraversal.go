package leetcode

import (
	"collections"
	"fmt"
)

func main() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2, Right: &node1}
	fmt.Println(preorderTraversal(&node2))
}

/**
二叉树的后序遍历，非递归
二叉树的后序遍历可以从前序遍历推出来
后序遍历：左 - 右  - 根
前序遍历：根 - 左  - 右
1、魔改前序遍历，先遍历右子树，在遍历左子树 变为  根 - 右  - 左
2、插入结果数组时使用头插法，变为 左 - 右 - 根
*/
func postorderTraversal(root *TreeNode) []int {
	stack := collections.NewStack()
	res := make([]int, 0)
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			res = append(res, node.Val)
			node = node.Right
		}
		node = stack.Pop().(*TreeNode)
		node = node.Left
	}
	// 反转数组
	return reverseArr(res)
}

func reverseArr(arr []int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
