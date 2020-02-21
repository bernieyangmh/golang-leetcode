package main

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


//看看别人系的
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