// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 示例 1：
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// 示例 2：
// 输入： 3
// 输出： 3
// 解释： 有三种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶 + 1 阶
// 2.  1 阶 + 2 阶
// 3.  2 阶 + 1 阶
package main

// 动态规划  dp的i代表前i个台阶的方法数, 是前面两个台阶方法数之和
// dp[i] = dp[i-2] + dp[i-1]
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]

}

// 只需要前两个值,压缩动态规划,
func climbStairs2(n int) int {

	// 1的时候直接返回
	if n == 1 {
		return 1
	}

	prev1 := 1
	prev2 := 1
	cnt := 0

	for i := 2; i <= n; i++ {
		cnt = prev1 + prev2
		prev1 = prev2
		prev2 = cnt
	}

	return cnt

}
