package main

//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
//示例 1:
//输入:
//Tree 1                     Tree 2
//1                         2
/// \                       / \
//3   2                     1   3
///                           \   \
//5                             4   7
//输出:
//合并后的树:
//3
/// \
//4   5
/// \   \
//5   4   7
//注意: 合并必须从两个树的根节点开始



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 自己写的，很慢很长
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	dfs(t1, t2)
	return t1
}

func dfs(t1 *TreeNode, t2 *TreeNode) {
	if t1.Left != nil && t2.Left != nil {
		t1.Left.Val += t2.Left.Val
	}
	if t1.Left == nil && t2.Left != nil {
		t1.Left = new(TreeNode)
		t1.Left.Val = t2.Left.Val
	}

	// t1.Left != nil && t2.Left == nil
	// t1.Left == nil && t2.Left == nil
	if t2.Left != nil {
		dfs(t1.Left, t2.Left)
	}

	if t1.Right != nil && t2.Right != nil {
		t1.Right.Val += t2.Right.Val
	}
	if t1.Right == nil && t2.Right != nil {
		t1.Right = new(TreeNode)
		t1.Right.Val = t2.Right.Val
	}

	// t1.Right != nil && t2.Right == nil
	// t1.Right == nil && t2.Right == nil
	if t2.Right != nil {
		dfs(t1.Right, t2.Right)
	}
	return
}



// 递归，很清晰
func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	res := new(TreeNode)
	res.Val = t1.Val+t2.Val
	res.Left = (mergeTrees(t1.Left, t2.Left))
	res.Right = mergeTrees(t1.Right, t2.Right)

	return res
}
