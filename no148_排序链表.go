package main

//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//示例 1:
//输入: 4->2->1->3
//输出: 1->2->3->4
//
//示例 2:
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var h, h1, h2, pre, res *ListNode

	h = head

	length, intv := 0, 1

	for h != nil {
		h = h.Next
		length++
	}

	res = new(ListNode)
	res.Val = 0
	res.Next = head

	for intv < length {
		pre = res
		h = res.Next
		for h != nil {
			i := intv
			h1 = h
			for i > 0 && h != nil {
				h = h.Next
				i--
			}
			if i > 0 {
				break
			}
			i = intv
			h2 = h
			for i > 0 && h != nil {
				h = h.Next
				i--
			}
			c1, c2 := intv, intv-1
			for c1 > 0 && c2 > 0 {
				if h2 == nil && h1.v || h1.Val < h2.Val {
					pre.Next = h1
					h1 = h1.Next
					c1--
				} else {
					pre.Next = h2
					h2 = h2.Next
					c2--
				}
				pre = pre.Next
			}
			if c1 == 0 {
				pre.Next = h2
			} else {
				pre.Next = h1
			}
			for c1 > 0 || c2 > 0 {
				pre = pre.Next
				c1--
				c2--
			}
			pre.Next = h
		}
		intv = intv * 2
	}
	return res.Next

}

func sortList2(head *ListNode) *ListNode {
	var (
		i, c1, c2 int
		pre, h1, h2 *ListNode
	)
	h, length, intv := head, 0, 1
	for h != nil {
		h, length = h.Next, length+1
	}
	res := &ListNode{Val: 0, Next: head}

	for intv < length {
		pre, h = res, res.Next
		for h != nil {
			h1, i = h, intv

			for i != 0 && h!= nil {
				h, i = h.Next, i-1
			}
			if i != 0 {
				break
			}

			h2, i = h, intv
			for i != 0 && h!=nil {
				h, i = h.Next, i-1
			}
			c1, c2 = intv, intv-1
			for c1 != 0 && c2 != 0 {
				if h1.Val < h2.Val{
					pre.Next, h1, c1 = h1, h1.Next, c1-1
				} else {
					pre.Next, h2, c2 = h2, h2.Next, c2-1
				}
				pre =pre.Next
			}
			if c1 != 0 {
				pre.Next = h1
			} else {
				pre.Next = h2
			}

			for c1 > 0 || c2 > 0 {
				pre, c1, c2 = pre.Next, c1-1, c2-1
			}
			pre.Next = h
		}
		intv = intv * 2
	}
	return res.Next

}
