package main

// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。
// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3 


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return enterLeaf(root, 1)
}

func enterLeaf(root *TreeNode, dep int) int {
	var depl, depr int
	depr = dep
	depl = dep
	if root.Left != nil {
		depl = enterLeaf(root.Left, dep+1)
	}
	if root.Right != nil {
		depr = enterLeaf(root.Right, dep+1)
	}
	return max(depl, depr)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
