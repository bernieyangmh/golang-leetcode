package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nln = new(ListNode)
	startNode := nln

	for 1==1 {
		if l1 != nil && l2 != nil{
			nln.Val += l1.Val + l2.Val

			l1 = l1.Next 
			l2 = l2.Next
			// 如果要进位或者后面仍有运算，新建一个节点
			if nln.Val >= 10 ||( l1 != nil || l2 != nil){
				nln.Next = new(ListNode)	
			}
			if nln.Val >= 10 {
				nln.Val = nln.Val %10
				nln.Next.Val += 1
			}
			if l1 == nil && l2 == nil{
				break
			}
			nln = nln.Next
			continue
		}
		if l1 == nil && l2 != nil {
			nln.Val += l2.Val
			l2 = l2.Next
			if nln.Val >= 10 ||( l1 != nil || l2 != nil){
				nln.Next = new(ListNode)	
			}
			if nln.Val >= 10 {
				nln.Val = nln.Val %10
				nln.Next.Val += 1
			}
			if l1 == nil && l2 == nil{
				break
			}
			nln = nln.Next
		}
		if l2 == nil && l1 != nil {
			nln.Val += l1.Val
			l1 = l1.Next
			if nln.Val >= 10 ||( l1 != nil || l2 != nil){
				nln.Next = new(ListNode)	
			}
			if nln.Val >= 10 {
				nln.Val = nln.Val %10
				nln.Next.Val += 1
			}
			if l1 == nil && l2 == nil{
				break
			}
			nln = nln.Next
		}

	}
return startNode
}


// Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }