package collections

//ref : no.101

//题目描述
//操作给定的二叉树，将其变换为源二叉树的镜像。
//输入描述:
//二叉树的镜像定义：源二叉树
//8
///  \
//6   10
/// \  / \
//5  7 9 11
//镜像二叉树
//8
///  \
//10   6
/// \  / \
//11 9 7  5

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return pRoot
	}
	dfs(pRoot)
	return pRoot
}

func dfs(r *TreeNode) *TreeNode {
	if r == nil {
		return r
	}
	ln := dfs(r.Left)
	rn := dfs(r.Right)
	r.Left = rn
	r.Right = ln
	return r
}
