// 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
// 说明：
// 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
// 示例 1:
// 输入: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// 输出: 1
// 示例 2:
// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历生成数组
func kthSmallest(root *TreeNode, k int) int {
	path := []int{}
	path = inOrder(root, path)
	return path[k-1]
}

func inOrder(root *TreeNode, path []int) []int {
	if root != nil {
		return path
	}
	path = inOrder(root.Left, path)
	path = append(path, root.Val)
	path = inOrder(root.Right, path)
	return path
}

// 迭代,用栈保存
func kthSmallest2(root *TreeNode, k int) int {
	var stack []*TreeNode
	for true {
		// 一直向左走
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 走到底,抛出最下面的左结点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		// 看看右边的结点路径
		root = root.Right
	}
	return 0 
}
