package main

//快指针1次走2步，慢指针1次走1步。所以快指针总是走了慢指针两倍的路。
//回顾一下阶段1的过程，设头节点到入环点的路途为 n, 那么慢指针走了入环路途的一半（n/2）时，快指针就到达入环点了(走完n了)。
//慢指针再继续走完剩下的一般入环路途（剩下的n/2），到达入环点时，快指针已经在环内又走了一个 n 那么远的路了。
//为了方便理解，这里先讨论环很大，大于n的情况（其他情况后文补充）。此时，慢指针正处于入环点，快指针距离入环点的距离为n。环内路，可以用此时快指针的位置分割为两段，前面的 n 部分，和后面的 b 部分。
//此时开始继续快慢指针跑圈，因为已经在环内了，他们其实就是在一条nbnbnbnbnbnbnb（无尽nb路）上跑步。
//慢指针从入环处开始跑b步，距离入环处就剩下了n。此时，快指针则是从距离入环处n步远的位置开始跑了2b步，距离入环处也是剩下了n。他们相遇了，并且距离入环处的距离就是n，n就是头节点到入环点的距离阿!!! 后面的不用说了吧。
//环很小的情况，其实跟环很大是一样的，比如你可以理解为将多个小环的循环铺开，虚拟扩展成一个大环来理解。


//作者：振之
//链接：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	//if head.Next.Next == nil || head.Next.Next

	fast, slow := head, head

	for {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			return nil
		}
		if fast == slow {
			break
		}
	}

	start := head
	for {
		start = start.Next
		slow = slow.Next
		if start == slow {
			return start
		}
	}
}