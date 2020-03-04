package main

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
//示例 :
//给定二叉树
//1
/// \
//2   3
/// \
//4   5
//返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//注意：两结点之间的路径长度是以它们之间边的数目表示

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res = 0

	dfs(root)
	return res - 1
}

func dfs(root *TreeNode) int {
	// return 0, 代表返回的是父节点左右子结点的相对深度，而不是总深度
	if root == nil {
		return 0
	}
	left := dfs(root.Right)
	right := dfs(root.Left)
	res = max(res, left+right+1)

	return max(left, right) + 1
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
