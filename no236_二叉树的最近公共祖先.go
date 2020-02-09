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

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	attr := 6
	fmt.Println(^(1 << 0))
	fmt.Println((^(1 << 0)) | (1 << 0))
	fmt.Println(attr&(^(1 << 0)) | (1 << 0))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		cur *TreeNode
	)
	_, cur = dfs(root, p, q, root)
	return cur
}

func dfs(root, p, q, cur *TreeNode) (bool, *TreeNode) {
	if root == nil {
		return false, root
	}
	var left, Right, mid int
	var lIS, rIS bool
	lIS, cur = dfs(root.Left, p, q, cur)
	rIS, cur = dfs(root.Right, p, q, cur)
	if lIS {
		left = 1
	} else {
		left = 0
	}

	if rIS {
		Right = 1
	} else {
		Right = 0
	}

	if root == p || root == q {
		mid = 1
	} else {
		mid = 0
	}
	if mid+left+Right >= 2 {
		cur = root
	}
	return mid+left+Right > 0, cur

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
		if hasP || hasQ {
			break
		}
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curNode.Left != nil {
			parentMap[root.Left] = root
			stack = append(stack, curNode.Left)
		}
		if curNode.Right != nil {
			parentMap[curNode.Right] = curNode
			stack = append(stack, curNode.Right)
		}

	}

	ancestors := make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parentMap[p]
	}
	for true {
		_, hasQ := ancestors[q]
		if hasQ {
			break
		}
		q = parentMap[q]
		return q
	}
	return q

}
