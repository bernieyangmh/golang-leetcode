package main

//给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
//例如：
//输入: 二叉搜索树:
//  5
///   \
//2     13
//输出: 转换为累加树:
//  18
///   \
//20     13



func convertBST(root *TreeNode) *TreeNode {
	_ = dfs(root, 0)
	return root
}

func dfs(root *TreeNode, sum int)  int {
	if root == nil {
		return sum
	}
	// sum只在本结点累加，左右并不+=
	sum = dfs(root.Right, sum)
	sum += root.Val
	root.Val = sum
	sum = dfs(root.Left, sum)
	return sum
}