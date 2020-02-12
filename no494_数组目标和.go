// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
// 示例 1:
// 输入: nums: [1, 1, 1, 1, 1], S: 3
// 输出: 5
// 解释:
// -1+1+1+1+1 = 3
// +1-1+1+1+1 = 3
// +1+1-1+1+1 = 3
// +1+1+1-1+1 = 3
// +1+1+1+1-1 = 3
// 一共有5种方法让最终目标和为3。
// 注意:
// 数组非空，且长度不会超过20。
// 初始的数组的和不会超过1000。
// 保证返回的最终结果能被32位整数存下。
package main

import "fmt"

func main() {
	a := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	S := 1
	fmt.Println(findTargetSumWays2(a, S))

}

// 递归,算每次相加减的结果递推到最后, 较慢
func findTargetSumWays(nums []int, S int) int {
	var res int

	res = dfs(nums, 0, S, 0, res)
	return res

}

func dfs(nums []int, point, S, sum, res int) int {
	if point == len(nums) {
		if sum == S {
			res += 1
		}
		return res
	}

	res = dfs(nums, point+1, S, sum+nums[point], res)
	res = dfs(nums, point+1, S, sum-nums[point], res)
	return res
}

// 前i个数和是j的组合数
// dp[i][j] = dp[i - 1][j - nums[i]] + dp[i - 1][j + nums[i]]
// 或
// dp[i][j + nums[i]] += dp[i - 1][j]
// dp[i][j - nums[i]] += dp[i - 1][j]

// j可能是负数,不支持负数作为下标, 1000以上为正,以下为负

func findTargetSumWays2(nums []int, S int) int {
	if S > 1000 {
		return 0
	}
	dp := [][]int{}
	for i := 0; i < len(nums); i++ {
		tmp := make([]int, 2001)
		dp = append(dp, tmp)
	}
	// 避免开头是0
	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000] += 1
	for i := 1; i < len(nums); i++ {
		for sum := -1000; sum <= 1000; sum++ {
			// 计算中有意义的sum会大于0
			if dp[i-1][sum+1000] > 0 {
				dp[i][sum+nums[i]+1000] += dp[i-1][sum+1000]
				dp[i][sum-nums[i]+1000] += dp[i-1][sum+1000]
			}
		}
	}
	return dp[len(nums)-1][S+1000]
}

// sum(P) - sum(N) = target
//		两边同时加上sum(P)+sum(N)）
// sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
// 		因为 sum(P) + sum(N) = sum(nums)
// 2 * sum(P) = target + sum(nums)
// sum(P)= [target+sum(nums)​]/2
// 作者：qsctech-sange
// 链接：https://leetcode-cn.com/problems/target-sum/solution/python-dfs-xiang-jie-by-jimmy00745/

// 一个只能取一次,满足和为p, 背包问题
func findTargetSumWays3(nums []int, S int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum < S || (sum+S)%2 == 1 {
		return 0
	}
	P := (sum + S) / 2

	dp := make([]int, 2001)
	dp[1000] = 1

	for _, n := range nums {
		for j := P; j >= n-1; j-- {
			dp[j+1000] += dp[j-n+1000]
		}
	}
	return dp[P+1000]
}
