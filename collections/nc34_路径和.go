package main


import (
	"fmt"
)

//一个机器人在m×n大小的地图的左上角（起点）。
//机器人每次向下或向右移动。机器人要到达地图的右下角（终点）。
//可以有多少种不同的路径从起点走到终点？

func main() {
	uniquePaths(2,2)
}

func uniquePaths(m int, n int) int {

	dp := [][]int{}
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		tmp[0] = 1
		dp = append(dp, tmp)
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 1
	}
	fmt.Println(dp)
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			// 如果能斜着走，就有dp[i-1][j-1]
			dp[i][j] = dp[i-1][j]+dp[i][j-1] // + dp[i-1][j-1]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}
