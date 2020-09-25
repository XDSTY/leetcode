package main

/**
Given a linked list, remove the n-th node from the end of list and return its head.
Example:
Given linked list: 1->2->3->4->5->nil, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:
Given n will always be valid.
Follow up:
Could you do this in one pass?
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	firstHead := head
	preHead := &ListNode{Next: head}
	for n > 0 && firstHead != nil {
		n--
		firstHead = firstHead.Next
	}
	for firstHead != nil {
		firstHead = firstHead.Next
		preHead = preHead.Next
	}
	if preHead.Next == head {
		if preHead.Next != nil {
			preHead.Next = preHead.Next.Next
		}
		return preHead.Next
	} else {
		if preHead.Next != nil {
			preHead.Next = preHead.Next.Next
		}
	}
	return head
}
