// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
// 你应当保留两个分区中每个节点的初始相对位置。
// 示例:
// 输入: head = 1->4->3->2->5->2, x = 3
// 输出: 1->2->2->4->3->5
package main

func main() {

}

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建两个哑结点,代表两个链表的头, 一个链表保存小于x的结点,一个保存大于等于x的结点
func partition(head *ListNode, x int) *ListNode {
	before_head := new(ListNode)
	before := before_head
	after_head := new(ListNode)
	after := after_head

	// head走到头
	for head != nil {
		// 放进before链表, before指针始终保持在before链表的最尾端
		// 放进after链表,before指针始终保持在before链表的最尾端
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	// after链表后面没结点,before后面是after链表,链接after链表的头
	after.Next = nil
	before.Next = after_head.Next
	// 返回before链表的头
	return before_head.Next

}
