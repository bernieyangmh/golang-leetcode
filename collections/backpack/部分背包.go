package backpack

//部分背包
//给定 n 种物品和一个容量为 C 的背包，物品 i 的重量是 wi，其价值为 vi，每个物品都有ki件，现在往背包里面装东西，怎么装能使背包的内物品价值最大？
//[][]int{0,1,2} 0是重量，1是价值，2是总数


func backpack03(nums [][]int, total int) int {
	dp := make([]int, total+1)

	// 容量为i时， 有多少个物品
	count := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		count = make([]int, total+1)
		for j := nums[i][0]; j <= total; j++ {
			if count[j-nums[i][0]] < nums[i][2] {
				dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])
				if dp[j] == dp[j-nums[i][0]]+nums[i][1] {
					count[j] = count[j-nums[i][0]] + 1
				}
			}
			if dp[j] == 0 || dp[j] < dp[j-1] {
				dp[j] = dp[j-1]
				count[j] = count[j-1]
			}
		}

	}
	return dp[total]
}
