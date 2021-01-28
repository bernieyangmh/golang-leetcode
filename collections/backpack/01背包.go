package backpack

//01背包问题

func backpack01(nums [][]int, total int) int {
	dp := make([][]int, len(nums))
	//初始化二维数组
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total+1)
	}

	//放入第一个物品，填第一行列表
	for i := nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1]
	}

	// 从后往前遍历可以省空间
	for i := 1; i < len(nums); i++ {
		for j := nums[i][0]; j < total; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i][0]]+nums[i][1])
		}
	}
	return dp[len(nums)-1][total]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
