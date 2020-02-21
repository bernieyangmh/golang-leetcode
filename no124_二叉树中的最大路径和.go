// 给定一个非空二叉树，返回其最大路径和。
// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
// 示例 1:
// 输入: [1,2,3]

//        1
//       / \
//      2   3

// 输出: 6
// 示例 2:
// 输入: [-10,9,20,null,null,15,7]
//    -10
//    / \
//   9  20
//     /  \
//    15   7

// 输出: 42

package main


var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = 1 << 31 * -1
	maxGain(root)
	return maxSum
}

//全局变量来更新
func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左边的路径最大值
	left_gain := max(maxGain(root.Left), 0)
	// 右边的路径最大值
	right_gain := max(maxGain(root.Right), 0)

	// 包含结点的路径
	newPath := root.Val + left_gain + right_gain

	maxSum = max(maxSum, newPath)

	// 对于该结点的父节点，只能走该结点的左边或右边
	return root.Val + max(left_gain, right_gain)

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
