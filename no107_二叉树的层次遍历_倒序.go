package main

// 笨方法，dfs最后倒序
func levelOrderBottom(root *TreeNode) [][]int {
	tmp := [][]int{}
	tmp = dfs(root, tmp, 0)
	ans := [][]int{}
	for i := len(tmp) - 1; i >= 0; i-- {
		ans = append(ans, tmp[i])
	}
	return ans
}

func dfs(root *TreeNode, res [][]int, level int) [][]int {
	if root == nil {
		return res
	}
	if len(res) <= level {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)
	if root.Left != nil {
		res = dfs(root.Left, res, level+1)

	}
	if root.Right != nil {
		res = dfs(root.Right, res, level+1)
	}

	return res
}

func levelOrderBottom2(root *TreeNode) [][]int {
	res := [][]int{}
	return dfs2(root, res, 0)
}

func dfs2(root *TreeNode, res [][]int, level int) [][]int {
	if root == nil {
		return res
	}
	if len(res) <= level {
		//res.insert(0, res)
		res = append([][]int{}, res...)
	}
	res = dfs(root.Left, res, level+1)
	res = dfs(root.Right, res, level+1)

	res[len(res)-level-1] = append(res[len(res)-level-1], root.Val)
	return res
}
