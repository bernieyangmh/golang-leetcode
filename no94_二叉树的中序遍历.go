// 给定一个二叉树，返回它的中序 遍历。
// 示例:
// 输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// 输出: [1,3,2]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历:左根右, 递归
func inorderTraversal(root *TreeNode) []int {
	var path []int
	path = dfs(root, path)

	return path
}

// 先加左边的,再加根,再加右边的
func dfs(root *TreeNode, path []int) []int {
	// 当前结点是空,退出递归
	if root == nil {
		return path
	}
	path = dfs(root.Left, path)
	path = append(path, root.Val)
	path = dfs(root.Right, path)
	return path
}

// 栈S;
// p = root;
// while(p || S不空){
//     while(p){
//         p入S;
//         p = p的左子树;
//     }
//     p = S.top 出栈;
//     访问p;
//     p = p的右子树;
// }
// 从p结点开始一直向左走,走到底返回,依次把路径上的值加到结果里,回到p,把p的值放到结果里,向右走
			
func inorderTraversal2(root *TreeNode) []int {
	var (
		res   []int
		stack []*TreeNode
	)
	p := root
	for p != nil && len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, p.Val)
		p = p.Right
	}

	return res
}
