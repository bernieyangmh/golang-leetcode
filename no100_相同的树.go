package main

//给定两个二叉树，编写一个函数来检验它们是否相同。
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p, q)
}

//1.都空 2.一个空 3.值不同 前中后序均可
func dfs(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !dfs(p.Left, q.Left) {
		return false
	}
	if !dfs(p.Right, q.Right) {
		return false
	}
	return true
}
