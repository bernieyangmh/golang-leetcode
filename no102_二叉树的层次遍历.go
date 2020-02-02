// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := "1321#123#412"
	fmt.Println(strconv.Atoi(""))
	fmt.Println(strings.Split(a, "#"))
}

// 递归,传递深度,根据深度判断当前结点应该在哪一层的数组里
// 数组没法index判断, 用字符串临时代替
func levelOrder(root *TreeNode) [][]int {
	var tmp []string
	var res [][]int
	if root == nil {
		return res
	}
	tmp = helper(root, 1, tmp)
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

func helper(root *TreeNode, depth int, tmp []string) []string {
	if root == nil {
		return tmp
	}
	if len(tmp) < depth {
		tmp = append(tmp, "")
	}
	tmp[depth-1] = tmp[depth-1] + "#" + strconv.Itoa(root.Val)
	tmp = helper(root.Left, depth+1, tmp)
	tmp = helper(root.Right, depth+1, tmp)
	return tmp
}
