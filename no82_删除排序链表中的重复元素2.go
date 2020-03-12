package main

//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//示例 1:
//输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//示例 2:
//输入: 1->1->1->2->3
//输出: 2->3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 加个伪结点,放进计算开始的条件,也方便返回值
	dummy := new(ListNode)
	dummy.Next = head
	flag := false
	pre := dummy

	for head.Next != nil {
		// 如果该结点的值和下一个结点的值相同,删除该结点,flag=true
		if head.Val == head.Next.Val {
			pre.Next = head.Next
			flag = true
		} else {
			// 如果flag=true,说明该结点也是重复的结点的,也删除,并把flag置为false
			if flag {
				pre.Next = head.Next
				flag = false
			} else {
				// 注意,之前删除结点的时候,pre是不同的,只有不删的时候pre正常移动
				pre = head
			}
		}
		// head移动到下一个,head删不删,next不变,所以正常移动就好
		head = head.Next
	}

	// 到底的时候也要检查一下标记, 111这种情况
	if flag {
		pre.Next = head.Next
	}
	return dummy.Next
}
