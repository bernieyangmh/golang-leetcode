// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
// 示例:
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL

package main

import (
	"fmt"
	"time"
)

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	s1 := new(ListNode)
	s2 := new(ListNode)
	s3 := new(ListNode)
	s4 := new(ListNode)
	s5 := new(ListNode)

	s1.Val = 1
	s1.Next = s2
	s2.Val = 2
	s2.Next = s3
	s3.Val = 3
	s3.Next = s4
	s4.Val = 4
	s4.Next = s5
	s5.Val = 5

	s := reverseBetween(s1, 1, 2)
	for s != nil {
		fmt.Print(s.Val)
		s = s.Next
		time.Sleep(time.Second)
	}

}

// 4指针迭代
// cur 指针当前的结点;prev指针 前一个结点; prev_tail 第一部分链表的尾; rever_tail第二部分链表的尾
//   第一部分       第二部分       第三部分
// 1 -> 1 -> | 2 -> 3 -> 4 | -> 5 -> 6 m=3 n=5
//      pt     rt=m     cur=n
// 1 -> 1 -> | 4 -> 3 -> 2 | -> 5 -> 6

// 如果 m > 1 新链表的头就是旧链表的头
// 如果 m = 1 新链表的头是第二部分链表的头
// | 1 -> 1 -> 2 | -> 3 -> 4 -> 5 -> 6 m=1 n=3
// | 2 -> 1 -> 1 | -> 3 -> 4 -> 5 -> 6 m=3 n=5
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n {
		return head
	}
	start := head
	prev := new(ListNode)
	prev_tail := new(ListNode)
	rever_tail := new(ListNode)

	var cur int
	for cur <= n {
		cur++
		// prev_tail 始终指向第一部分有序链表的最尾, prev及时更新,head前进
		if cur < m {
			prev = head
			prev_tail = head
			head = head.Next
			continue

		}

		// prev_tail 应该指向的第二部分的最尾还没出现,在此之后prev_tail固定不变
		// rever_tail 反转链表尾应该指向的第三部分的头还没出现,先记下
		if cur == m {
			rever_tail = head
			prev = head
			head = head.Next
			continue
		}
		// 1.记下当前指针的下一个,方便之后继续前进
		// 2.当前指针的下一个是上一个
		// 3.上一个指针前进
		// 4.当前指针前进
		if m < cur && cur < n {
			tmp := head.Next
			head.Next = prev
			prev = head
			head = tmp
			continue
		}
		// 走到了反转链表的尾端
		// 第一部分的尾指向当前指针,第二部分反转链表的头
		// 反转链表的尾指向下一个结点,第三部分的头
		// 当前指针的结点指向前一个
		if cur == n {
			rever_tail.Next = head.Next
			prev_tail.Next = head
			head.Next = prev
			break
		}
	}
	// 如果m是1,说明第二部分整体在前,新链表的头是第二部分的头,即循环结束cur指向的结点
	if m == 1 {
		start = head
	}
	return start
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	prev := new(ListNode)
	for m > 1 {
		prev = cur
		cur = cur.Next
		m--
		n--
	}
	con := prev
	tail := cur
	for n > 0 {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
		n--
	}
	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	tail.Next = cur
	return head
}
