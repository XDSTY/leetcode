package leetcode

import "fmt"

/**
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
思路：使用queue，queue的队首保存当前合法的最大值，当新的值大于队首的值时，可以清空队列，该值入队
	新的值小于队首时，在递增队列中找到自己合适的位置
*/

func main() {
	nums := []int{9, 10, 9, -7, -4, -8, 2, -6}
	arr := maxSlidingWindow(nums, 5)
	fmt.Println(arr)
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || len(nums) == 0 {
		return []int{}
	}
	res := make([]int, len(nums)-k+1)
	queue := New()
	for i := 0; i < len(nums); i++ {
		if queue.head != nil && queue.peek() < i-k+1 {
			queue.poll()
		}
		if queue.head == nil {
			queue.offer(i)
		} else if nums[queue.peek()] <= nums[i] {
			// 队列的头元素（最大）小于当前元素，清空队列
			queue.reset()
			queue.offer(i)
		} else {
			// 从队列的头部开始找到自己的位置
			head := queue.head
			for head.next != nil && nums[head.next.val] > nums[i] {
				head = head.next
			}
			head.next = &Node{val: i, next: nil}
			queue.tail = head.next
		}
		if i-k+1 >= 0 {
			res[i-k+1] = nums[queue.peek()]
		}
	}
	return res
}

type (
	Queue struct {
		head *Node
		tail *Node
	}
	Node struct {
		next *Node
		val  int
	}
)

func New() *Queue {
	return &Queue{nil, nil}
}

func (queue *Queue) offer(data int) {
	if queue.head == nil {
		queue.head = &Node{val: data, next: nil}
		queue.tail = queue.head
	} else {
		queue.tail.next = &Node{val: data, next: nil}
		queue.tail = queue.tail.next
	}
}

func (queue *Queue) poll() int {
	if queue == nil || queue.head == nil {
		return -1
	}
	node := queue.head
	queue.head = node.next
	return node.val
}

func (queue *Queue) peek() int {
	if queue == nil || queue.head == nil {
		return -1
	}
	return queue.head.val
}

func (queue *Queue) reset() {
	queue.head = nil
	queue.tail = nil
}
