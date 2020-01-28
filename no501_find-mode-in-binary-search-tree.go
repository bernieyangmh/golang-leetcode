package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var currentArr []int
	var maxCnt, num int
	num = 1
	var pre *TreeNode
	inOrder(root, pre, num, maxCnt, currentArr)
	return currentArr
}

func inOrder(root, pre *TreeNode, num, maxCnt int, currentArr []int) {
	if root == nil {
		return
	}
	inOrder(root.Left, pre, num, maxCnt, currentArr)
	if pre != nil {
		if pre.Val == root.Val {
			num += 1
		} else {
			num = 1
		}
	}

	if num == maxCnt {
		currentArr = append(currentArr, root.Val)
	}
	if num > maxCnt {
		currentArr = []int{}
		currentArr = append(currentArr, root.Val)
		maxCnt = num
	}
	pre = root
	inOrder(root.Right, pre, num, maxCnt, currentArr)
}

1. num:1 maxcnt:1 arr[1]
2.num:1 maxcnt: 1 arr [1, 2]
