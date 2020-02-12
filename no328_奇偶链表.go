// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
// 示例 1:
// 输入: 1->2->3->4->5->NULL
// 输出: 1->3->5->2->4->NULL
// 示例 2:
// 输入: 2->1->3->5->6->4->7->NULL
// 输出: 2->3->6->7->1->5->4->NULL
// 说明:
// 应当保持奇数节点和偶数节点的相对顺序。
// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

package main

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 像拉链一样拆分链表
// 记录下奇链表的尾和偶链表的尾,和当前结点是奇偶的flag, 根据flag决定当前结点加到哪个链表的尾.最后缝合两个链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 从第三个开始
	cur := head.Next.Next
	flag := 1

	odd_tail := head
	even_tail := head.Next

	odd_head := head
	even_head := head.Next

	for cur != nil {
		// 当前结点是奇数
		if flag == 1 {
			odd_tail.Next = cur
			odd_tail = cur
			flag = 2
		} else {
			even_tail.Next = cur
			even_tail = cur
			flag = 1
		}
		cur = cur.Next
	}

	odd_tail.Next = even_head
	// 
	even_tail.Next = nil
	return odd_head
}
