package main

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//说明: 叶子节点是指没有子节点的节点。
//示例:
//给定二叉树 [3,9,20,null,null,15,7],
//3
/// \
//9  20
///  \
//15   7
//返回它的最小深度  2.

var minDepthNum int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDepthNum = 1 << 32
	dfs(root, 0)

	return minDepthNum
}

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		minDepthNum = min(minDepthNum, depth+1)
		return
	}
	dfs(root.Left, depth+1)
	dfs(root.Right, depth+1)
}























func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


