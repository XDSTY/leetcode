package leetcode

/**
Given a linked list, swap every two adjacent nodes and return its head.
You may not modify the values in the list's nodes, only nodes itself may be changed.
交换相邻的节点
Example:
Given 1->2->3->4, you should return the list as 2->1->4->3.
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preHead := &ListNode{Next: head}
	preNode := preHead
	currNode := head
	var nextNode *ListNode
	for currNode != nil && currNode.Next != nil {
		preNode.Next = currNode.Next
		preNode = currNode
		nextNode = currNode.Next.Next
		currNode.Next.Next = currNode
		currNode.Next = nextNode
		currNode = currNode.Next
	}
	return preHead.Next
}

func swapPairsRecr(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		headNode := head.Next
		nextNode := headNode.Next
		head.Next.Next = head
		head.Next = swapPairs(nextNode)
		return headNode
	}
	return head
}
