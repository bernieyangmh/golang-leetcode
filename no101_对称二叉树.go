// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//     1
//    / \
//   2   2
//    \   \
//    3    3
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	// 两边都是空,可以
	if t1 == nil && t2 == nil {
		return true
	}
	// 一边有值一遍空不行
	if t1 == nil || t2 == nil {
		return false
	}
	// 两个结点要相等,他们的镜像结点要相等 左边的左子树要等与右边的右子树, 左边的右子树要等于右边的左子树
	return t1.Val == t2.Val && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)

}

// 递归
// 每次取出两个结点判断
// 每次放入4个结点
func isSymmetric2(root *TreeNode) bool {
	stack := []*TreeNode{}
	stack = append(stack, root)
	stack = append(stack, root)

	for len(stack) > 0 {
		t1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		t2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}
		stack = append(stack, t1.Left)
		stack = append(stack, t2.Right)
		stack = append(stack, t1.Right)
		stack = append(stack, t2.Left)

	}
	return true

}
