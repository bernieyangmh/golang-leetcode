package main

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//示例 1:
//输入: 1->1->2
//输出: 1->2
//示例 2:
//输入: 1->1->2->3->3
//输出: 1->2->3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy

	for head.Next != nil {
		if head.Val == head.Next.Val {
			pre.Next = head.Next
		} else {
			pre = head
		}
		head = head.Next
	}
	return dummy.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	cursor := head
	for cursor != nil && cursor.Next != nil {
		next := cursor.Next
		if next.Val == cursor.Val {
			cursor.Next = next.Next
		} else {
			cursor = cursor.Next
		}
	}
	return head
}
