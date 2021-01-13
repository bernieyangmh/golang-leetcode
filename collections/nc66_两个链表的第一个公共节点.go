package collections

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	//两个指针， 走完一个链表开始走另一个
	// 如果有公共节点。p1=p2的时候跳出循环，没有就都走到尾的nil
	p1, p2 := pHead1, pHead2
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 != p2 {
			if p1 == nil {
				p1 = pHead2
			}
			if p2 == nil {
				p2 = pHead1
			}
		}
	}

	return p1

}
