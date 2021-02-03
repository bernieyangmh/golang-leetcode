package main

// 反转链表 n~m部分

import (
	"fmt"
)

func main() {
	n1 := new(ListNode)
	n2 := new(ListNode)
	n3 := new(ListNode)
	n4 := new(ListNode)
	n5 := new(ListNode)
	n6 := new(ListNode)
	n7 := new(ListNode)

	n1.Next = n2
	n1.Val = 1
	n2.Next = n3
	n2.Val = 2
	n3.Next = n4
	n3.Val = 3
	n4.Next = n5
	n4.Val = 4
	n5.Next = n6
	n5.Val = 5
	n6.Next = n7
	n6.Val = 6
	n7.Next = nil
	n7.Val = 7
	n := reverPartListNode(n1, 7, 1)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转m～n的链表 m > n
func reverPartListNode(root *ListNode, m, n int) *ListNode {
	//双指针确定要翻转的部分链表的头尾

	var dummy1, dummy2 *ListNode
	dummy := new(ListNode)
	dummy.Next = root

	fast := dummy
	slow := dummy

	// fast 走M步，slow 走N步， 确定要反转的链表的首尾
	for m > 0 {
		fast = fast.Next
		m--
	}
	// 记录下fast的下一个，翻转后的tail 要指向该node, 是nil也没关系
	dummy2 = fast.Next

	for n > 1 {
		// n暂时少走一步， 这个dummy之后要链接 返回的head
		slow = slow.Next
		n--
	}
	// dummy1要用来连接返回的head
	dummy1 = slow
	slow = slow.Next

	//12 345 67 -> 12 543 67
	//5是head， 3是tail， 2是dummy1， 6是dummy2
	head, tail := reverse(slow, fast)
	// 3 -> 6
	tail.Next = dummy2
	//m > n slow不会是nil 2->5
	dummy1.Next = head
	return dummy.Next
}

// 反转给定的链表， 返回链表的首尾
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		// 先记录下一个node
		next := p.Next
		//将当前节点的指向改为前一个
		p.Next = prev
		// 双指针均向后移动一位
		prev = p
		p = next
	}
	// 尾是头， 头是尾
	return tail, head
}
