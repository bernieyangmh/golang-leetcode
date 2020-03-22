package main

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Val = -1<<31
	pre, tail := dummy, dummy
	cur := head

	for cur != nil {
		if tail.Val < cur.Val {
			tail.Next = cur
			tail =cur
			cur = cur.Next
		} else {
			tmp := cur.Next
			tail.Next = tmp
			for pre.Next != nil && pre.Next.Val < cur.Val {
				pre = pre.Next
			}
			cur.Next = pre.Next
			pre.Next = cur
			pre = dummy
			cur = tmp
		}
	}
	return dummy.Next
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	vNodeMap := make(map[int]*ListNode)
	arrValue := []int{head.Val}
	vNodeMap[head.Val] = head
	dummy := head.Next

	for dummy != nil {
		cur := len(arrValue)-1
		for cur >= 0 && arrValue[cur] >= dummy.Val {
			cur--
		}
		// dummy就是最小的
		if cur == -1 {
			arrValue = append([]int{dummy.Val}, arrValue...)
			vNodeMap[dummy.Val] = dummy
			tmp := dummy.Next
			dummy.Next = vNodeMap[arrValue[0]]
			dummy = tmp

		} else {
			arrValue = append(arrValue, 0)   // Step 1
			copy(arrValue[cur:], arrValue[1:]) // Step 2
			arrValue[cur] = dummy.Val
			vNodeMap[dummy.Val] = dummy

			prev := vNodeMap[arrValue[cur]]
			tmp := dummy.Next
			dummy.Next = prev.Next
			prev.Next = dummy
			dummy = tmp
		}
	}
	return vNodeMap[arrValue[0]]
}