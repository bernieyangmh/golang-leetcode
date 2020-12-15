package main

import (
	"fmt"
)

//完全二叉树的定义
//若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树广度搜索， 因为都集中在左边，所以最后才会是nil
func isCompleteTree(root *TreeNode) bool {
	var tailIsNil bool // 最后一个节点是否是nil
	
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			tailIsNil = true
		} else {
			if tailIsNil {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}
