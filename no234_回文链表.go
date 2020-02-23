package main

import (
	"fmt"
)

//请判断一个链表是否为回文链表。
//示例 1:
//输入: 1->2
//输出: false
//示例 2:
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？




// 遍历一遍，数组记录，然后数组从右向左，指针从链表左向右
func isPalindrome(head *ListNode) bool {
	arr := []int{}
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	if len(arr) < 2 {
		return true
	}

	i:=0
	for i<len(arr)/2 {
		if head.Val != arr[len(arr)-1-i] {
			return false
		}
		i++
		head = head.Next
	}
	return true
}


//	一个指针走两步，一个指针走一步，当前者走到底时，后者在中间(奇)或中间左边(偶)
// 	一个指针从头走到慢指针，期间反转链表，然后一个指针从中间从左到右，一个从右到左，判断是否回文
	// 把无结点，一个结点，两个结点的情况先判断，因为我们先判断.Next.Next再移动
func isPalindrome2(head *ListNode) bool {

	if head == nil || head.Next == nil{
		return true
	}
	if head.Next.Next == nil {
		if head.Next.Val != head.Val {
			return false
		} else {
			return true
		}
	}

	isOdd := false
	oneS, towS, reverPoint := head, head, head
	var dummy *ListNode

	for towS != nil {
		// odd
		// oneS is mid
		if towS.Next == nil {
			isOdd = true
			break
		}
		// even
		// oneS is left part
		if towS.Next.Next == nil {
			break
		}
		oneS = oneS.Next
		towS = towS.Next.Next
	}

	// reverse left part list
	for true {

		after := reverPoint.Next
		reverPoint.Next = dummy
		dummy = reverPoint
		reverPoint = after

		// dummy是已经被反转的链表，如果用reverPoint的话，还需做一次反转操作
		if dummy == oneS {
			// rever，oneS分别代表原链表的右部分和左部分
			if isOdd {
				// 奇数的话，最中间的数我们不判断，因为next只能指向一个
				oneS = dummy.Next
			}
			break
		}
	}

	for oneS!=nil && reverPoint != nil {
		if oneS.Val != reverPoint.Val {
			return false
		}
		oneS = oneS.Next
		reverPoint = reverPoint.Next
	}
	return true
}

// 别人的反转, 不用快慢指针而是计算
func isPalindrome3(head *ListNode) bool {
	n := 0
	for p:=head; p!=nil; p=p.Next {
		n ++
	}
	// 找到链表中点的下一个位置
	p := head
	for i:=0; i<n/2; i++ {
		p = p.Next
	}
	if n%2 == 1 {
		p = p.Next
	}
	// 翻转后半段链表
	p = reverse(p)
	for p != nil {
		if head.Val != p.Val {
			return false
		}
		p = p.Next
		head = head.Next
	}
	return true
}
// 反转一个链表，并返回反转后到头节点
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}










var frontPointer *ListNode

func isPalindrome4(head *ListNode) bool {
	frontPointer = head
	return  recurCheck(head)
}

func recurCheck(cur *ListNode) bool {
	if cur != nil {
		if !recurCheck(cur.Next) {
			return false
		}
		if cur.Val != frontPointer.Val {
			return false
		}
		frontPointer = frontPointer.Next
	}
	return true

}