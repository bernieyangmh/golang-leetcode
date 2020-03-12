package main

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右

//动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	dp := [][]int{}

	//初始化第一行，遇到障碍，之后的路径为0
	num := 1
	for i:=0;i<m;i++{
		tmp := make([]int, n)
		dp = append(dp, tmp)
		if obstacleGrid[i][0] == 1 {
			num=0
			//这里是continue，仍需继续添加行。。。
			continue
		}
		dp[i][0] = num
	}
	// 初始化第一列
	num = 1
	for i:=0;i<n;i++{
		if obstacleGrid[0][i] == 1 {
			num = 0
			break
		}
		dp[0][i] = num
	}

	// 遇到障碍，该路径的dp就是0
	// 如果不是障碍，到达该位置的路径的数量等于上面加左边
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//返回最后一位的数量
	return dp[m-1][n-1]
}