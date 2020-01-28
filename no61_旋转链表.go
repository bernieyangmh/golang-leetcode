// 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
// 示例 1:
// 输入: 1->2->3->4->5->NULL, k = 2
// 输出: 4->5->1->2->3->NULL
// 解释:
// 向右旋转 1 步: 5->1->2->3->4->NULL
// 向右旋转 2 步: 4->5->1->2->3->NULL
// 示例 2:
// 输入: 0->1->2->NULL, k = 4
// 输出: 2->0->1->NULL
// 解释:
// 向右旋转 1 步: 2->0->1->NULL
// 向右旋转 2 步: 1->2->0->NULL
// 向右旋转 3 步: 0->1->2->NULL
// 向右旋转 4 步: 2->0->1->NULL
package main

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var one = new(ListNode)
	one.Val = 0
	one.Next = nil
	rotateRight(nil, 1)
	for one != nil {
		println(one.Val)
		one = one.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var n = 1
	var old_tail, new_tail, new_head *ListNode
	old_tail = head
	// 累加计算多少个结点
	// 判断到尾结点
	for old_tail.Next != nil {
		n += 1
		old_tail = old_tail.Next
	}
	// 拼接链表
	old_tail.Next = head
	new_head = head
	// 指向新的尾结点
	// [1,2,3,4,5]n=5, k=2, n-(k%n)=3 指针走了3步,到达3 [4,5,1,2,3]
	// [0,1,2]n=3, k=4,n-(k%n)=2, 指针走了2步,到达1 [2,0,1]
	for i := 0; i < n-(k%n); i++ {
		// 记录尾
		new_tail = new_head
		// 更新头
		new_head = new_head.Next
	}
	new_tail.Next = nil
	return new_head

}
