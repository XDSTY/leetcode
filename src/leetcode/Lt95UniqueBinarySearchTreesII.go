package leetcode

import "fmt"

/**
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
给一个整数n，返回组成的二叉搜索树
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
/**
划分为求左右子树的问题
*/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	var getTrees func(start, end int) []*TreeNode
	getTrees = func(start, end int) []*TreeNode {
		allTrees := make([]*TreeNode, 0)
		if start > end {
			allTrees = append(allTrees, nil)
			return allTrees
		}
		for i := start; i <= end; i++ {
			leftTrees := getTrees(start, i-1)
			rightTrees := getTrees(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					root := &TreeNode{Val: i, Left: left, Right: right}
					allTrees = append(allTrees, root)
				}
			}
		}
		return allTrees
	}
	return getTrees(1, n)
}

func main() {
	allTrees := make([]*TreeNode, 0)
	allTrees = append(allTrees, nil)
	for _, val := range allTrees {
		fmt.Println(val)
	}
}
