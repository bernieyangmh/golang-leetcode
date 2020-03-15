package main

import (
	"fmt"
)

//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//示例 1:
//给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//示例 2:
//给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main()  {
	prev := &ListNode{Val:1}
	p := prev
	for i:=2;i<8;i++{
		cur := &ListNode{Val:i}
		prev.Next = cur
		prev = cur
	}
	reorderList3(p)
	//
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

}

//线性表暴力
func reorderList(head *ListNode) {
	arr := []*ListNode{}
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	if len(arr) < 3 {
		return
	}
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left].Next = arr[right]
		arr[right].Next = arr[left+1]
		left++
		right--
	}
	if left == right {
		arr[left-1].Next = arr[left]
	}
	return
}

//递归
func reorderList2(head *ListNode) {
	length := 0
	dummy := head
	for dummy != nil {
		dummy = dummy.Next
		length++
	}
	if length <3 {
		return
	}
	helper(head, length)
}

//1,2,3,4,5,6,7,8,9,10
func helper(head *ListNode, length int) *ListNode {
	// 处理之后把尾返回，1是奇数情况，2是偶数情况
	if length == 1 {
		fmt.Println(head)
		tail := head.Next
		head.Next = nil
		return tail
	}

	if length == 2 {
		//7
		tail := head.Next.Next
		//6->nil
		head.Next.Next = nil
		return tail
	}
	tail := helper(head.Next, length-2)
	//5
	subHead := head.Next
	//4->7
	head.Next = tail
	//8
	outail := tail.Next
	//7 -> 5
	tail.Next = subHead
	return outail
}

// 翻转
func reorderList3(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	isEven := false
	// 快慢指针找中点
	fast := head
	mid := head
	for fast.Next != nil && fast.Next.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		isEven = true
	}
	//翻转中点后的链表，如果是偶数的话，中点属于前一个链表
	var pre *ListNode
	cur := mid.Next
	for cur!=nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	// 拼接两个链表
	//1234765
	//mid.Next = pre
	//fmt.Println(head, mid)
	////从两个链表头开始合并
	//p1 := head
	//p2 := mid.Next
	//// 奇数mid只有最后一个 1726354
	//// 偶数mid和mid之后的不用动 162534
	//for p1 != mid {
	//	mid.Next = p2.Next
	//	p2.Next = p1.Next
	//	p1.Next = p2
	//	p1 = p2.Next
	//	p2 = mid.Next
	//}
	p1 := head
	p2 := pre
	if isEven {
		mid.Next.Next = nil
	} else {
		mid.Next = nil
	}

	fmt.Println(p1, p2, isEven, mid.Next)
	for p1 != mid {
		nextp1 := p1.Next
		nextp2 := p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = nextp1
		p2 = nextp2

	}

}















