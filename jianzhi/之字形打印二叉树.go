package main

import (
	"fmt"
)

func main() {
	// q1    1
	// q2  2   3
	// q1 4 5 6 7
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(PrintNode(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintNode(pRoot *TreeNode) [][]int {
	// write code here
	res := [][]int{}
	if pRoot == nil {
		return res
	}

	var stack1, stack2 []*TreeNode
	
	stack1 = append(stack1, pRoot)
	for len(stack1) > 0 || len(stack2) > 0 {
		var res1 []int
		var res2 []int

		for len(stack1) > 0 {
			cur := stack1[len(stack1)-1]

			if cur.Left != nil {
				stack2 = append(stack2, cur.Left)
			}
			if cur.Right != nil {
				stack2 = append(stack2, cur.Right)
			}
			res1 = append(res1, cur.Val)
			stack1 = stack1[:len(stack1)-1]
		}
		if len(res1)>0 {
			res = append(res, res1)
		}
		for len(stack2) > 0 {
			cur := stack2[len(stack2)-1]
			if cur.Right != nil {
				stack1 = append(stack1, cur.Right)
			}
			if cur.Left != nil {
				stack1 = append(stack1, cur.Left)
			}
			res2 = append(res2, cur.Val)
			stack2 = stack2[:len(stack2)-1]
		}
		if len(res2) > 0 {
			res = append(res, res2)
		}
	}
	return res
}
