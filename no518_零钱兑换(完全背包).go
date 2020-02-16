package main

func main() {

}

// 剪枝,超时间限制
func change(amount int, coins []int) int {
	return dfs(amount, coins, 0)

}

func dfs(target int, coins []int, res int) int {
	if target == 0 {
		return res + 1
	}
	for i, v := range coins {
		if v <= target {
			// fmt.Println(v, target, res)
			// 只允许拿之后的硬币,不然会出现1+2+2 || 2+2+1的情况
			// 根据这一点就有动态规划的思路
			res = dfs(target-v, coins[i:], res)
		}
	}
	return res
}

// 动态规划 dp代表金额构成的组合数
func change2(amount int, coins []int) int {
	size := len(coins)
	if size == 0 {
		if amount == 0 {
			return 1
		}
		return 0
	}
	dp := make([]int, amount+1)
	
	// 初始化
	//用于j-coins[i]==0时 一个硬币的情况
	dp[0] = 1
	
	// 一个硬币组成的全部金额
	for i := coins[0]; i <= amount; i += coins[0] {
		dp[i] = 1
	}
	
	// i代表使用前i个硬币, j代表组成的金额 数值代表组合数
	for i := 1; i < size; i++ {
		// 从当前硬币金额的面值到总面值
		for j := coins[i]; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}

	return dp[amount]
}
//i/j 0  1  2  3  4  5
// 1 1, 1, 1, 1, 1, 1
// 2 1, 1, 2, 2, 3, 3
// 5 1, 1, 2, 2, 3, 4

