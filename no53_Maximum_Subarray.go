// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

package main

import (
	"fmt"
)

func main() {
	nnn := []int{1, 213, 5, 23423, 13, 12}
	fmt.Println(maxSubArray(nnn))

}

// 我们定义函数 S(i) ，它的功能是计算以 0（包括 0）开始加到 i（包括 i）的值。
// 那么 S(j) - S(i - 1) 就等于 从 i 开始（包括 i）加到 j（包括 j）的值。
// 我们进一步分析，实际上我们只需要遍历一次计算出所有的 S(i), 其中 i = 0,1,2....,n-1。
// 然后我们再减去之前的 S(k),其中 k = 0，1，i - 1，中的最小值即可。
// 因此我们需要 用一个变量来维护这个最小值，还需要一个变量维护最大值。
func maxSubArray(nums []int) int {
	var maxValue, minsum, sum int
	if len(nums) == 0 {
		return 0
	}

	//初始化最大值，可能最大值比0还低
	maxValue = nums[0]
	for _, i := range nums {
		sum += i
		// 先判断最大值，在改最小值
		if sum-minsum >= maxValue {
			maxValue = sum - minsum
		}
		if sum <= minsum {
			minsum = sum
		}
	}
	return maxValue
}

// 动态规划
// dp[i] - 表示到当前位置 i 的最大子序列和
// 状态转移方程为： dp[i] = max(dp[i - 1] + nums[i], nums[i])
// 初始化：dp[0] = nums[0]
// 从状态转移方程中，我们只关注前一个状态的值，所以不需要开一个数组记录位置所有子序列和，只需要两个变量，
// currMaxSum - 累计最大和到当前位置i
// maxSum - 全局最大子序列和:
// currMaxSum = max(currMaxSum + nums[i], nums[i])
// maxSum = max(currMaxSum, maxSum)
func dpmaxSubArray(nums []int) int {
	var currMaxSum, maxSum int
	currMaxSum = nums[0]
	maxSum = nums[0]
	for index, i := range nums {
		if index == 0 {
			continue
		}
		currMaxSum = max(currMaxSum+i, i)
		maxSum = max(currMaxSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
