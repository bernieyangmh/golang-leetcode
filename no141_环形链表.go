package main


//给定一个链表，判断链表中是否有环。
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//如果链表中存在环，则返回 true 。 否则，返回 false 。

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	first := head
	second := head

	for first != nil {
		if first.Next == nil {
			return false
		}
		first = first.Next
		if second.Next == nil || second.Next.Next == nil {
			return false
		}
		second = second.Next.Next
		if first == second {
			return true
		}
	}
	return false
}


//看看别人写的
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	first := head
	second := head.Next

	// 如果没环，快指针先到头，判断快指针是否是nil就可以
	for first != second {
		if second == nil || second.Next == nil {
			return false
		}
		first = first.Next
		second = second.Next.Next
	}
	return true
}