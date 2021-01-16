package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {
	// write code here
	stack := []int{}
	if head == nil {
		return true
	}

	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	for len(stack)/2 > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if head.Val != top {
			return false
		}
		head = head.Next
	}
	return true
}
