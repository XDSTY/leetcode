package leetcode

/**
Given the head of a linked list, return the list after sorting it in ascending order.
对链表进行排序
Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
在时间复杂度O(n logn)
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list

链表的归并或者快速排序
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var getMidNode func(node *ListNode) *ListNode
	var mergeSort func(node *ListNode) *ListNode
	getMidNode = func(node *ListNode) *ListNode {
		slow, fast := node, node.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		res := slow.Next
		slow.Next = nil
		return res
	}
	mergeSort = func(node *ListNode) *ListNode {
		if node == nil || node.Next == nil {
			return node
		}
		midNode := getMidNode(node)
		ln := mergeSort(node)
		rn := mergeSort(midNode)

		preNode := &ListNode{}
		tmpNode := preNode
		for ln != nil && rn != nil {
			if ln.Val < rn.Val {
				tmpNode.Next = ln
				ln = ln.Next
			} else {
				tmpNode.Next = rn
				rn = rn.Next
			}
			tmpNode = tmpNode.Next
		}
		if ln != nil {
			tmpNode.Next = ln
		}
		if rn != nil {
			tmpNode.Next = rn
		}
		return preNode.Next
	}
	return mergeSort(head)
}
