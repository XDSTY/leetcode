package niuke

func ReverseList(pHead *ListNode) *ListNode {
	var preNode *ListNode
	curr := pHead
	var next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = preNode
		preNode = curr
		curr = next
	}
	return preNode
}
