package main



func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum==0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}