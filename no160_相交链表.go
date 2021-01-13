package main


//编写一个程序，找到两个单链表相交的起始节点。
//
//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
//输出：Reference of the node with value = 8
//输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。


// 1.暴力，对于a上的每个结点，遍历b看有没有
// 2.map，保存a上的所有结点，看b上面有没有相同的

// 3.双指针，一个从a的头后移，一个从b的头后移；走到尾时分别记录下尾结点
// a指针走到尾，则从b链表开始走；b指针走到尾，则从a链表开始走，直到两个指针指向相同的结点，该结点是相交起始结点
// 两个链表的尾不一样，说明两个链表不相交
// 起点虽然不一样，但路程一样a+b,速度一样，必定同时到达
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA, pointB := headA, headB

	// 记录两个链表的尾部
	var tailA, tailB *ListNode

	// 排除空链表情况
	if pointA == nil || pointB == nil {
		return nil
	}

	// 一直走到两个指针指向相同结点
	// 除非两个尾结点不相同
	for pointA != pointB {

		// 走到尾，标记A的尾节点，从B再走
		if pointA.Next == nil {
			tailA = pointA
			pointA = headB
		} else {
			pointA = pointA.Next
		}

		if pointB.Next == nil {
			tailB = pointB
			pointB = headA
		} else {
			pointB = pointB.Next
		}

		// 判断尾结点
		if tailA != nil && tailB != nil {
			if tailA != tailB {
				return nil
			}
		}

	}
	return pointA
}