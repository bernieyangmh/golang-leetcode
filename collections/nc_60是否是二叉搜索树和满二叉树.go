package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a1 := &TreeNode{Val: 2}
	a2 := &TreeNode{Val: 1}
	a3 := &TreeNode{Val: 3}
	a1.Left = a2
	a1.Right = a3

	fmt.Println(judgeIt(a1))

}

/**
 *
 * @param root TreeNode类 the root
 * @return bool布尔型一维数组
 */

func judgeIt(root *TreeNode) []bool {
	stack := []*TreeNode{}
	var inorder int
	var flag = false
	var is_ss = true
	var is_full bool
	var count float64

	// 走到底且处理完 栈全部值,代表走到最右边结点
	for root != nil || len(stack) > 0 {
		for root != nil {
			// push
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		// 第一次判断一定成立,用flag=false 代表inorder此时是无限小,任何值都能小于它, 不可能满足该条件
		if root.Val <= inorder && flag {
			if is_ss {
				is_ss = false
			}
		}
		//  2
		// 1 3
		// 走到1到底,回退到结点1, 此时flag是false, inorder 是1, 右结点是nil, 再从栈拿出2
		// 2>1, inorder值更新为2, 进入右结点
		inorder = root.Val
		flag = true
		root = root.Right
	}

	for i := 0.0; i < 100; i++ {
		if math.Pow(2.0, i)-1 == count {
			is_full = true
			break
		}
	}
	return []bool{is_ss, is_full}
}
