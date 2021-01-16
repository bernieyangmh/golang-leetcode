// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。 树的右视图
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

// 先进入左结点再进入右结点,右结点的值覆盖左结点
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

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	help(root, &result, 1)
	return result
}

func help(root *TreeNode, result *[]int, i int) {
	if len(*result) < i {
		*result = append(*result, root.Val)
	}
	// 先递归右，直到最大层数
	if root.Right != nil {
		help(root.Right, result, i+1)
	}

	// 左边层数比右高，再用左
	if root.Left != nil {
		help(root.Left, result, i+1)
	}
	return
}
