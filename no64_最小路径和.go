// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例:
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。

package main

func minPathSum(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}
	n := len(grid[0])
	m := len(grid)

	dp := [][]int{}
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}
			//i--m--第一列
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			//j--n--第一行
			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 把上面的dp全换成grid, 在原数组上修改
func minPathSum2(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				grid[0][0] = grid[0][0]
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
				continue
			}
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	M, N := len(grid), len(grid[0])
	dp := make([][]int, M)
	for i := 0; i < M; i++ {
		dp[i] = make([]int, N)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < M; i++ {
		dp[i][0] = dp[i-1][0]+grid[i][0]
	}
	for j := 1; j < N; j++ {
		dp[0][j] = dp[0][j-1]+grid[0][j]
	}
	for x := 1; x < M; x++ {
		for y := 1; y < N; y++ {
			dp[x][y] = min(dp[x-1][y], dp[x][y-1]) + grid[x][y]
		}
	}
	return dp[M-1][N-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
