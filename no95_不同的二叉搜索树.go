// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
// 示例:
// 输入: 3
// 输出:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// 解释:
// 以上的输出对应以下 5 种不同结构的二叉搜索树：

//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateTrees(3))

}

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从序列 1...n 中取出数字i,作为当前树的树根。于是，剩余 i - 1 个元素可用于左子树，n - i 个元素用于右子树。
// 这样会产生 G(i - 1) 种左子树 和 G(n - i) 种右子树，其中 G 是卡特兰数。
// 现在，我们对序列 1 ... i - 1 重复上述过程，以构建所有的左子树；然后对 i + 1 ... n 重复，以构建所有的右子树。
// 这样，我们就有了树根 i 和可能的左子树、右子树的列表。
// 最后一步，对两个列表循环，将左子树和右子树连接在根上
func generateTrees(n int) []*TreeNode {
	var treeArr = []*TreeNode{}
	if n == 0 {
		return treeArr
	}
	return generate_tree(1, n)
}

func generate_tree(s, e int) []*TreeNode {
	allTree := []*TreeNode{}
	if s > e {
		allTree = append(allTree, nil)
		return allTree
	}
	// 遍历一遍数组
	// 左边的组成左子树数组, 右边的组成右子树数组, 该结点和每个左右子树都拼一下
	for i := s; i <= e; i++ {
		leftTree := generate_tree(s, i-1)
		rightTree := generate_tree(i+1, e)
		for _, lt := range leftTree {
			for _, rt := range rightTree {
				curTree := new(TreeNode)
				curTree.Val = i
				curTree.Left = lt
				curTree.Right = rt
				allTree = append(allTree, curTree)
			}

		}
	}
	return allTree
}
