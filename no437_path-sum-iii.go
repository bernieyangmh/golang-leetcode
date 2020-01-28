// 给定一个二叉树，它的每个结点都存放着一个整数值。
// 找出路径和等于给定数值的路径总数。
// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
// 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
// 示例：
// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//       10
//      /  \
//     5   -3
//    / \    \
//   3   2   11
//  / \   \
// 3  -2   1
// 返回 3。和等于 8 的路径有:
// 1.  5 -> 3
// 2.  5 -> 2 -> 1
// 3.  -3 -> 11

package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 以每个节点为根节点，都算一遍路径和为sum的有几条
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// 以每个节点为根节点
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum -= root.Val
	if sum == 0 {
		return 1 + helper(root.Left, sum) + helper(root.Right, sum)
	} else {
		return 0 + helper(root.Left, sum) + helper(root.Right, sum)
	}
}

// 以每个节点为根节点，都算一遍路径和为sum的有几条
func pathSum2(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var hashmap = make(map[int]int)
	// 以每个节点为根节点
	return helper2(root, sum, hashmap, 0)
}

func helper2(root *TreeNode, sum int, hm map[int]int, p int) int {
	if root == nil {
		return 0
	}
	var res int
	p += root.Val

	if p == sum {
		res += 1
	}

	if value, ok := hm[p-sum]; ok {
		res += value
	}
	if _, ok := hm[p]; ok {
		hm[p] += 1
	} else {
		hm[p] = 1
	}
	res = helper2(root.Left, sum, hm, p) + helper2(root.Right, sum, hm, p) + res
	hm[p] = hm[p] + 1
	return res
}
