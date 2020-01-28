// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 示例：
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
// 给定的 n 保证是有效的。
// 进阶：
// 你能尝试使用一趟扫描实现吗？

package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 超时
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var start = head
	var point = head
	var num int
	for head.Next != nil {
		if num >= 2 {
			point = point.Next
		}
		num += 1
		head = head.Next
	}
	point.Next = head
	return start
}

// 善于利用哑结点
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := dummy

	// 先走几步
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	// [1,2,3,4,5], first 走到5时前移到nil, second移动到3
	for first != nil {
		first = first.Next
		second = second.Next
	}
	// 删除seconde前面的结点
	second.Next = second.Next.Next
	return dummy.Next
}
