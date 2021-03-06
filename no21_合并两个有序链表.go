// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例：
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy_head := new(ListNode)
	cur := dummy_head

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 都没走到头
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	// 直接相连
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummy_head.Next
}
