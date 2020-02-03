// 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
// 例如，从根到叶子节点路径 1->2->3 代表数字 123。
// 计算从根到叶子节点生成的所有数字之和。
// 说明: 叶子节点是指没有子节点的节点。
// 示例:
// 输入: [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// 输出: 1026
// 解释:
// 从根到叶子节点路径 4->9->5 代表数字 495.
// 从根到叶子节点路径 4->9->1 代表数字 491.
// 从根到叶子节点路径 4->0 代表数字 40.
// 因此，数字总和 = 495 + 491 + 40 = 1026
package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0, 0)
}

func helper(root *TreeNode, res int, pathNum int) int {
	if root == nil {
		return res
	}
	// 当前值放到路径值后面
	pathNum = pathNum*10 + root.Val

	// 走到叶子结点,结果+
	if root.Left == nil && root.Right == nil {
		res += pathNum
		return res
	}
	// 因为res不是全局的,每次更新都要重新赋值
	res = helper(root.Left, res, pathNum)
	res = helper(root.Right, res, pathNum)
	return res
}
