package leetcode

/**
Merge two sorted linked lists and return it as a new sorted list.
The new list should be made by splicing together the nodes of the first two lists.
合并两个有序的链表
Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	tmpNode := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmpNode.Next = l1
			l1 = l1.Next
		} else {
			tmpNode.Next = l2
			l2 = l2.Next
		}
		tmpNode = tmpNode.Next
	}
	if l1 != nil {
		tmpNode.Next = l1
	}
	if l2 != nil {
		tmpNode.Next = l2
	}
	return head.Next
}
