// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 示例:
// 给定 1->2->3->4, 你应该返回 2->1->4->3.

package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
// 123456, 每次递归进入函数,seconde先指向first,first指向下一层函数的second,
// 2->1, 4->3, 6->5; 函数到底head=nil后收敛; 5->nil, 3->6, 1->4
// 2-1-4-3-6-5-nil
// 12345; 2->1 4->3 3->5; 2-1-4-3-5
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next

	first.Next = swapPairs(second.Next)
	second.Next = first
	return second
}

// 迭代,每次3个结点参与
func swapPairs2(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	prev := dummy
	// first 和second 每次迭代时根据head初始化,方便边界计算
	// 1234, prev:1,first:2,second:3; 23要交换; 
	// 1->3 2->4 3->2 1324 
	for head != nil && head.Next != nil {
		first := head
		second := head.Next
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next
	}
	return dummy.Next
}
