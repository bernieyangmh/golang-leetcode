// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
// 示例 1:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
// 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
// 示例 2:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出: 5
// 解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
// 说明:
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉树中。

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力法 map记录每个结点从根结点到该点的路径
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pathMap := make(map[int][]*TreeNode)
	path := []*TreeNode{}

	pathMap = dfs(root, path, pathMap)
	pPath := pathMap[p.Val]
	qPath := pathMap[q.Val]
	n := 0
	for true {
		if n == len(pPath) || n == len(qPath) {
			return pPath[n-1]
		}
		if pPath[n] != qPath[n] {
			return pPath[n-1]
		}

		n++
	}
	return root

}

func dfs(root *TreeNode, path []*TreeNode, pathMap map[int][]*TreeNode) (map[int][]*TreeNode) {
	if root == nil {
		return pathMap
	}

	path = append(path, root)
	pathMap[root.Val] = append([]*TreeNode{}, path...)

	pathMap = dfs(root.Left, path, pathMap)
	pathMap = dfs(root.Right, path, pathMap)
	return pathMap

}

// 递归 left,right, mid 分别代表p，q在左子树或右子树或就在该结点
// 三者满足两个，就说明该结点是p,q的公共祖先， 该结点再往上，就不会>=2
var res *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	_ = dfs1(root, p, q)
	return res
}

func dfs1(root, p, q *TreeNode) (bool) {
	var left, right, mid int
	if root == nil {
		return false
	}
	if dfs1(root.Left, p, q) {
		left = 1
	} else {
		left = 0
	}

	if dfs1(root.Right, p, q) {
		right = 1
	} else {
		right = 0
	}

	if root == p || root == q {
		mid = 1
	} else {
		mid = 0
	}

	if mid+left+right >= 2 {
		res = root
	}
	// 找到了p或q才可能大于0
	return mid+left+right > 0

}

// 父指针迭代
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	parentMap := make(map[*TreeNode]*TreeNode)
	parentMap[root] = nil

	var hasP, hasQ bool
	for true {
		_, hasP = parentMap[p]
		_, hasQ = parentMap[q]
		if hasP && hasQ {
			break
		} else {
			curNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if curNode.Left != nil {
				parentMap[root.Left] = curNode
				stack = append(stack, curNode.Left)
			}
			if curNode.Right != nil {
				parentMap[curNode.Right] = curNode
				stack = append(stack, curNode.Right)
			}
		}
	}

	ancestors := make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parentMap[p]
	}
	for true {
		_, hasQ := ancestors[q]
		if !hasQ {
			q = parentMap[q]
		} else {
			break
		}

	}
	return q

}
