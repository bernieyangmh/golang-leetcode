package main

// 暴力法
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func isBalanced2(root *TreeNode) bool {
	is, _ := helper(root)
	return is
}

func helper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, -1
	}
	lb, lh := helper(root.Left)
	if !lb {
		return false, 0
	}
	rb, rh := helper(root.Right)
	if !rb {
		return false, 0
	}
	return abs(lh-rh) < 2, 1+max(lh, rh)
}

func isBalanced2(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	// 走到底，不改变高度
	if root == nil {
		return 0
	}
	// 左边不平衡就可以直接返回
	left := dfs(root.Left)
	if left == -1 {
		return -1
	}
	// 右边不平衡就可以直接返回
	right := dfs(root.Right)
	if right == -1 {
		return -1
	}

	不平衡
	if abs(left-right) > 1 {
		return -1
	}
	//返回当前距底的高度
	return max(left, right) + 1
}