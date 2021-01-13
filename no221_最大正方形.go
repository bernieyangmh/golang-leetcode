// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
// 示例:
// 输入:
// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 1 1 1
// 1 0 0 1 0
// 输出: 4
package main

import "fmt"

// 0 1 1 1 0	0 1 1 1 0
// 1 1 1 1 0	1 1 2 2 0
// 0 1 1 1 1 ->	0 1 2 3 1
// 0 1 1 1 1	0 1 2 3 2
// 0 0 1 1 1	0 0 1 2 3

// 动态规划的思路, dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
// 当前的数字如果是1, dp在该位置的值是左边,上边,左上3个值的最小值 在+1
// 当前的数字如果是0, 保持不变

// 有点像扫雷
// 1 1
// 1 ? -> ? 这个该是几

// 1 1 2
// 1 2 2
// 2 2 ? -> 这个该是几

// [["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]]
func main() {
	a := "01"
	fmt.Println(a[0], a[1])

}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var dp [][]int
	var maxEdge int
	m := len(matrix)

	for i := 0; i < m; i++ {
		dp = append(dp, []int{})
		for j := 0; j < m; j++ {
			dp[i] = append(dp[i], 0)
			if i == 0 || j == 0 {
				if matrix[i][j] == 48 {
					dp[i][j] = 0
				} else {
					maxEdge = max(maxEdge, 1)
					dp[i][j] = 1
				}
				continue
			}
			if matrix[i][j] == 48 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
			maxEdge = max(maxEdge, dp[i][j])
		}
		fmt.Println(dp)
	}
	return maxEdge * maxEdge

}

// 优化了一下
func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var dp [][]int
	var maxEdge int

	for i := 0; i < len(matrix); i++ {
		tmp := make([]int, len(matrix[0]))
		dp = append(dp, tmp)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// 是0，包含这个点的无法组成正方形
			if matrix[i][j] == 48 {
				dp[i][j] = 0
				continue
			}
			if maxEdge == 0 {
				maxEdge = 1
			}
			// 初始化 第一行、第一列
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			//这个点 dp在该位置的值是左边,上边,左上3个值的最小值 再+1 包含这个点
			dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
			maxEdge = max(maxEdge, dp[i][j])
		}
	}
	return maxEdge * maxEdge
}

// 一维够了,二维数组性能也高
//        | prev是1 
// dp   0 1 2 3 2
//    	0 0 1 2 3 
//          | 走到1时
func maximalSquare3(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxEdge, prev int

	// 多一位,方便计算,开头是没有左上角的
	dp := make([]int, len(matrix[0])+1)

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			tmp := dp[j]
			if matrix[i-1][j-1] == 49 {
				// j-i是左边,prev是左上,j是上边
				dp[j] = min(min(dp[j-1], prev), dp[j]) + 1
				maxEdge = max(maxEdge, dp[j])
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}
	return maxEdge * maxEdge
}

// 位操作的骚操作
// https://leetcode-cn.com/problems/maximal-square/solution/xing-yu-xing-zhi-jian-jin-xing-yu-cao-zuo-zai-ji-s/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
