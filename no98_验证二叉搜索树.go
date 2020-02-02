// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
// 假设一个二叉搜索树具有如下特征：
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:
// 输入:
//     2
//    / \
//   1   3
// // 输出: true
// 示例 2:
// 输入:
//     5
//    / \
//   1   4
//      / \
//     3   6
// 输出: false
// 解释: 输入为: [5,1,4,null,null,3,6]。
//      根节点的值为 5 ，但是其右子节点值为 4 。
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 设置边界,每次递归更新边界值
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 第一次从根结点进入无边界条件的限制
	return helper(root, root.Val, root.Val, 0, 0)

}

func helper(root *TreeNode, lower, upper, islower, isupper int) bool {
	if root == nil {
		return true
	}

	if islower != 0 && root.Val <= lower {
		return false
	}
	if isupper != 0 && root.Val >= upper {
		return false
	}
	// 进入左结点,左结点的值在最小值和根结点之间
	if !helper(root.Left, lower, root.Val, islower, 1) {
		return false
	}
	// 进入右结点,右结点的值在最大值和根结点之间
	if !helper(root.Right, root.Val, upper, 1, isupper) {
		return false
	}

	return true

}

// 中序遍历,左根右应该是从大到小的顺序排列
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	var inorder int
	var flag = false

	// 走到第且处理完 栈全部值,代表走到最右边结点
	for root != nil || len(stack) > 0 {
		for root != nil {
			// push
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 第一次判断一定成立,用flag=false 代表inorder此时是无限小,任何值都能小于它, 不可能满足该条件
		if root.Val <= inorder && flag {
			return false
		}
		//  2
		// 1 3
		// 走到1到底,回退到结点1, 此时flag是false, inorder 是1, 右结点是nil, 再从栈拿出2
		// 2>1, inorder值更新为2, 进入右结点
		inorder = root.Val
		flag = true
		root = root.Right
	}
	return true
}
