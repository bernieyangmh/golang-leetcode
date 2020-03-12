package main

//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//示例:
//给定的有序链表： [-10, -3, 0, 5, 9],
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//0
/// \
//-3   9
///   /
//-10  5

// 二分
func sortedListToBST(head *ListNode) *TreeNode {
	length := findSize(head)
	dummy, _ := convertToBST(0, length, head)
	return dummy
}

func findSize(head *ListNode) int {
	pre := head
	l := 0
	for pre != nil {
		pre = pre.Next
		l++
	}
	return l
}

func convertToBST(left, right int, node *ListNode) (*TreeNode, *ListNode) {
	if left >= right {
		return nil, node
	}
	mid := (left + right) / 2
	LeftRoot, node := convertToBST(left, mid, node)
	root := new(TreeNode)
	root.Val = node.Val
	root.Left = LeftRoot
	node = node.Next

	RightRoot, node := convertToBST(mid+1, right, node)
	root.Right = RightRoot

	return root, node
}
