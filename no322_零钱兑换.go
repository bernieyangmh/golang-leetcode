// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
// 示例 1:
// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3
// 解释: 11 = 5 + 5 + 1
// 示例 2:
// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:
// 你可以认为每种硬币的数量是无限的。
package main

import "fmt"

func main() {
	a := []int{186, 419, 83, 408}
	fmt.Println(coinChange(a, 1024))

}

// 状态转移方程
// 0 n=0,
// 1+min{f(n-Ci)} i=[1,k]
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)

	// 硬币没有小于1的,所以总个数不可能比自身的值还大,
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	//初始化状态
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 不同面额来比较
		for _, v := range coins {
			// 硬币面额比金额还大的不考虑
			if i >= v {
				dp[i] = min(dp[i], dp[i-v]+1)

			}
		}
	}

	//比总金额还大的说明没有修改,走不到该点
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//暴力递归
func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := 1 << 32
	for _, c := range coins {
		if amount-c < 0 {
			continue
		}
		subCnt := coinChange2(coins, amount-c)
		if subCnt == -1 {
			continue
		}
		res = min(res, subCnt+1)
	}
	if res == 1<<32 {
		return -1
	}
	return 0
}

// 自顶向下递归, 边界问题
func coinChange3(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -2
	}
	return dfs(coins, memo, amount)
}

func dfs(coins, memo []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if memo[amount] != -2 {
		return memo[amount]
	}
	var res = 1 << 32
	for _, c := range coins {
		if amount-c < 0 {
			continue
		}
		subCnt := dfs(coins, memo, amount-c)
		if subCnt == -1 {
			continue
		}
		res = min(res, subCnt+1)
	}
	if memo[amount] == 1<<32 {
		return -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}
