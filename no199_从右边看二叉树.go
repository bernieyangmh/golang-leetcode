// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// 示例:
// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:
// 1            <---
// /   \
// 2     3         <---
// \     \
// 5     4       <---

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先进入左结点再进入有结点,右结点的值覆盖左结点
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	// 初始化根结点的值
	res = append(res, root.Val)
	if root.Right == nil && root.Left == nil {
		return res
	}
	res = helper(root, res, 1)
	return res
}

func helper(root *TreeNode, res []int, depth int) []int {
	// 没有叶子结点,返回
	if root.Right == nil && root.Left == nil {
		return res
	}
	// 很重要,同一层只需一个结点新增res,不然会有多个0的情况发生
	if len(res) == depth {
		res = append(res, 0)
	}
	// 左边有,这一层的结果是左
	if root.Left != nil {
		res[depth] = root.Left.Val
		res = helper(root.Left, res, depth+1)
	}
	// 右边有,这一层的结果属于右
	if root.Right != nil {
		res[depth] = root.Right.Val
		res = helper(root.Right, res, depth+1)
	}

	return res
}
