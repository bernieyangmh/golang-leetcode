// 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
// 说明: 叶子节点是指没有子节点的节点。
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// 返回:
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力,走到每个叶子结点, 判断路径和是否满足
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	helper(root, &res, []int{}, sum, 0)
	return res
}

func helper(root *TreeNode, res *[][]int, path []int, sum, accumulation int) {
	fmt.Println(root, path, accumulation)
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Val+accumulation == sum && (root.Left == nil && root.Right == nil) {
		tmp := append([]int(nil), path...)
		*res = append(*res, tmp)
		return
	}
	helper(root.Left, res, path, sum, accumulation+root.Val)
	helper(root.Right, res, path, sum, accumulation+root.Val)
}

func pathSum2(root *TreeNode, sum int) (rst [][]int) {
	pathSum113(root, sum, []int{}, &rst)
	return
}

func pathSum113(root *TreeNode, sum int, paths []int, rst *[][]int) {
	if root == nil {
		return
	}
	paths = append(paths, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			// 这样深copy 貌似更快点, leetcode后台不稳
			var tmp = make([]int, len(paths))
			copy(tmp, paths)
			*rst = append(*rst, tmp)
			return
		}
	}
	if root.Left != nil {
		pathSum113(root.Left, sum-root.Val, paths, rst)
	}
	if root.Right != nil {
		pathSum113(root.Right, sum-root.Val, paths, rst)
	}
}
