// 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回锯齿形层次遍历如下：
// [
//   [3],
//   [20,9],
//   [15,7]
// ]

package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 同102, 多了判断走向的flag
func zigzagLevelOrder(root *TreeNode) [][]int {
	var tmp []string
	var res [][]int
	if root == nil {
		return res
	}
	tmp = helper(root, 1, tmp, true)
	for _, value := range tmp {
		arr := strings.Split(value, "#")
		intArr := []int{}
		for _, s := range arr {
			i, err := strconv.Atoi(s)
			if err == nil {
				intArr = append(intArr, i)
			}
		}
		res = append(res, intArr)
	}

	return res
}

func helper(root *TreeNode, depth int, tmp []string, leftFirst bool) []string {
	if root == nil {
		return tmp
	}
	if len(tmp) < depth {
		tmp = append(tmp, "")
	}
	// 如果从左向右,每次就在右边追加
	// 否则,从左边追加
	if leftFirst {
		tmp[depth-1] = tmp[depth-1] + "#" + strconv.Itoa(root.Val)

	} else {
		tmp[depth-1] = strconv.Itoa(root.Val) + "#" + tmp[depth-1]
	}
	tmp = helper(root.Left, depth+1, tmp, !leftFirst)
	tmp = helper(root.Right, depth+1, tmp, !leftFirst)

	return tmp
}
