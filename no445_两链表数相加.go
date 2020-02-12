// 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
// 进阶:
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
// 示例:
// 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出: 7 -> 8 -> 0 -> 7

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 用双栈存放两个链表的值,遍历完链表之后一一弹出拼接,注意边界值和细节
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var (
		l1Stack, l2Stack          []*ListNode
		l1Length, l2Length, carry int
		// 两个指针方便操作
		cur, prev *ListNode
	)

	for l1 != nil {
		l1Stack = append(l1Stack, l1)
		l1Length++
		l1 = l1.Next
	}

	for l2 != nil {
		l2Stack = append(l2Stack, l2)
		l2Length++
		l2 = l2.Next
	}

	// 两个链表都没走到头
	for l1Length != 0 || l2Length != 0 {
		cur = new(ListNode)
		// 先处理prev
		cur.Next = prev
		// 如果两个栈都有
		if l1Length != 0 && l2Length != 0 {
			cur.Val = l1Stack[l1Length-1].Val + l2Stack[l2Length-1].Val + carry
			l1Length--
			l2Length--
		} else if l1Length != 0 {
			cur.Val = l1Stack[l1Length-1].Val + carry
			l1Length--
		} else if l2Length != 0 {
			cur.Val = l2Stack[l2Length-1].Val + carry
			l2Length--
		}
		// 更新进位, 余数是值
		carry = cur.Val / 10
		cur.Val = cur.Val % 10
		prev = cur
	}

	// 最后的进位如果有,处理一次
	if carry > 0 {
		prev = cur
		cur = new(ListNode)
		cur.Val = carry
		cur.Next = prev
	}
	return cur
}

// 链表的的反转
func reverse(head *ListNode) *ListNode {

	var prev *ListNode
	cur := head
	after := cur.Next
	for cur != nil {
		after = cur.Next
		cur.Next = prev
		prev = cur
		cur = after
	}
	return prev
}
