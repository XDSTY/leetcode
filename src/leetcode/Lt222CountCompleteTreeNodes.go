package leetcode

import "fmt"

/**
给出一个完全二叉树，求出该树的节点个数。
说明：
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
示例:
输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-complete-tree-nodes

思路：可以对底层的节点进行编号，假设node.left编号0，node.right编号1，则底层最左边的编号为0000，右边的node编号为0001，最右边的编号为1111
可以通过二分法确定底层节点的个数
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var getLevelNode func(node *TreeNode, val int, level uint32) bool
	getLevelNode = func(node *TreeNode, val int, level uint32) bool {
		bit := 1 << (level - 1)
		for bit > 0 && node != nil {
			if (bit & val) == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bit >>= 1
		}
		return node != nil
	}

	var level uint32
	level = 0
	node := root
	for node.Left != nil {
		node = node.Left
		level++
	}
	left := 1 << level
	right := (1 << (level + 1)) - 1

	for left < right {
		mid := (left + right + 1) / 2
		// 判断mid是否是为nil，不为nil，则left = mid + 1，否则right = mid - 1
		if getLevelNode(root, mid, level) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3, Left: n6}
	n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	fmt.Println(countNodes(n1))

	fmt.Println(7 & 2)
}
