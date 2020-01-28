// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

package main

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummyHead = new(ListNode)
		p         = l1
		q         = l2
		curr      = dummyHead
		carry     = 0
		x         = 0
		y         = 0
		sum       = 0
	)
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum = carry + x + y
		carry = sum / 10
		// 赋值的是下一个结点的值, 如果pq都为nil,下一个结点是nil
		curr.Next = new(ListNode)
		curr.Next.Val = sum % 10
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = carry
	}
	return dummyHead.Next
}
