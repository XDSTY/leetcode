package leetcode

import (
	"container/heap"
)

/**
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.
合并k个有序链表
Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:
Input: lists = []
Output: []
Example 3:
Input: lists = [[]]
Output: []
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
思路：使用k个链表维护一个最小堆
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LN []*ListNode

func (heap *LN) Len() int           { return len(*heap) }
func (heap *LN) Less(i, j int) bool { return (*heap)[i].Val < (*heap)[j].Val }
func (heap *LN) Swap(i, j int)      { (*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i] }

func (heap *LN) Push(x interface{}) {
	*heap = append(*heap, x.(*ListNode))
}

func (heap *LN) Pop() interface{} {
	val := (*heap)[heap.Len()-1]
	*heap = (*heap)[0 : heap.Len()-1]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	newList := make([]*ListNode, 0)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			newList = append(newList, lists[i])
		}
	}
	h := LN(newList)
	heap.Init(&h)
	preHead := &ListNode{}
	tmpHead := preHead
	for h.Len() != 0 {
		node := heap.Pop(&h).(*ListNode)
		tmpHead.Next = node
		tmpHead = tmpHead.Next
		if node != nil && node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}
	return preHead.Next
}
