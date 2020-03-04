package main

//No.198 打家劫舍
//No.213 首尾相邻
//No.337 二叉树形式

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
// 示例 1:
// 输入: [1,2,3,1]
// 输出: 4
// 解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 2:
// 输入: [2,7,9,3,1]
// 输出: 12
// 解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//      偷窃到的最高金额 = 2 + 9 + 1 = 12 。

// 动态规划
// 我们本质上在解决对于第[i] 个房子，我们抢还是不抢。的问题。
// 判断的标准就是总价值哪个更大， 那么对于抢的话就是当前的房子可以抢的价值 + dp[i - 2]
// i - 1 不能抢，否则会触发警铃
// 如果不抢的话，就是dp[i - 1].
// 状态转移方程 dp[i] = Math.max(dp[i - 2] + nums[i - 2], dp[i - 1])

// 1.自顶向下 -- 生成一个全是-1的数组，然后递归，更新数组，最后返回
// 2.自底向上 -- 生成一个l+2数组，然后从后向前遍历, 返回dp[0]
// 3.自底向上O(1) 因为数组只用到了前两天的数据，可以只有两个变量
func rob198(nums []int) int {
	// 前两天，前一天，现在
	var dp_2, dp_1, dp_i int
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		dp_i = max(dp_1, dp_2+nums[i])
		dp_2 = dp_1
		dp_1 = dp_i
	}
	return dp_i

}

// 与No.198相比，我们要考虑三种情况
//1.首尾房子都不抢 (长度短，被2，3包含，忽略)
//2.只抢首
//3.只抢尾

func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	return max(No198Change(nums, 0, len(nums)-2), No198Change(nums, 1, len(nums)-1))
}

func No198Change(nums []int, start, end int) int {
	var dp_2, dp_1, dp_i int

	// 这里是小于等于，我们传的是可以访问的索引
	for i := start; i <= end; i++ {
		dp_i = max(dp_1, dp_2+nums[i])
		dp_2 = dp_1
		dp_1 = dp_i
	}
	return dp_i

}

func rob337(root *TreeNode) int {
	res := dfs(root)
	return max(res[0],res[1])
}

// 0不拿，1拿
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0,0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)

	noRobNum := max(left[0], left[1]) + max(right[0], right[1])
	robNum := root.Val+left[0]+right[0]
	return []int{noRobNum, robNum}
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
