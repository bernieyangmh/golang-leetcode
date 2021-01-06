// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 示例:
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

package main

import (
	"fmt"
)

// 两次遍历,
func maxSubArray(nums []int) int {
	// 初始化最大值是最小值
	maxNum := 1 << 31 * -1
	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			// 每加一个更新一次最大值
			maxNum = max(maxNum, tmp)
		}
	}
	return maxNum
}

// dp, 前i个字符的最大连续和
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		// 状态转移方程  要么是他自身,要么自身加前一个最大和
		// 因为是前一个,可以优化为O(1)
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	maxSum := 1 << 32 * -1
	for i := 0; i < len(dp); i++ {
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

// dp压缩成O1
func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prev := nums[0]
	maxSum := prev

	for i := 1; i < len(nums); i++ {
		// 状态转移方程  要么是他自身,要么自身加前一个最大和
		// 因为是前一个,可以优化为O(1)
		prev = max(nums[i], prev+nums[i])
		maxSum = max(maxSum, prev)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 最大和的非连续子数组
func maxNoSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		tmp := dp[1]
		// 1要当前数，0+cur
		dp[1] = dp[0] + nums[i]
		// 0 不要当前数，前一个或前两个
		dp[0] = max(dp[0], tmp)
	}
	return max(dp[0], dp[1])
}
