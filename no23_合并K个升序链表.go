package main

import (
	"container/heap"
)

//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//示例 1：
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//输入：lists = []
//输出：[]
//示例 3：
//输入：lists = [[]]
//输出：[]

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// 和No.21一样
func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	// 一个虚假的节点， 两个链表不断比较，加入
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)
	return mergeTwoSortedLists(left, right)
}

//构建一个最小堆
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy

	if len(lists) == 0 {
		return dummy.Next
	}

	pq := make(PQ, 0)
	for i, _ := range lists {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		// 获取K个链表的最小头
		item := heap.Pop(&pq).(*ListNode)

		// 给新链表加节点
		p.Next = item
		p = p.Next

		// 将next节点放入最小堆中
		if item.Next != nil {
			heap.Push(&pq, item.Next)
		}
	}
	return dummy.Next
}

type PQ []*ListNode

func (p PQ) Len() int { return len(p) }
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}
