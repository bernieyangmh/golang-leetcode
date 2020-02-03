// 给定一个二叉树，返回它的 前序 遍历。

//  示例:

// 输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// 输出: [1,2,3]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根左右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return preorder(root, []int{})
}

func preorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preorder(root.Left, res)
	res = preorder(root.Right, res)
	return res
}

// 利用一个栈,每次抛出最上面的结点,进入结点后,根值放进结果,右结点先进栈左结点再进
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return res
}
