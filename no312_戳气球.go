package main

//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 
//这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
//求所能获得硬币的最大数量。
//
//说明:
//你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
//0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
//示例:
//输入: [3,1,5,8]
//输出: 167
//解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

//https://leetcode-cn.com/problems/burst-balloons/solution/chao-xiang-xi-hui-su-dao-fen-zhi-dao-dp-by-niu-you/

var maxCoin int

// 暴力，超时 O(!N)
//第一层的所有例子
//[-1,1,5,8]
//[3,-1,5,8]
//[3,1,-1,8]
//[3,1,5,-1]

func maxCoins(nums []int) int {
	// 这里归零，leetcode所有测试用例都用这一个变量
	maxCoin = 0
	dfs(nums, 0, len(nums), 0)
	return maxCoin
}

func dfs(nums []int, s, length, beforeCoins int) {
	// 如果走到头，更新
	if s == length {
		// 这一次扎破的顺序产生的积大于最大值
		if beforeCoins > maxCoin {
			maxCoin = beforeCoins
		}
		return
	}

	for i := 0; i < length; i++ {
		// -1代表已经走过
		if nums[i] == -1 {
			continue
		}
		//保留该值，方便回溯
		temp := nums[i]
		nums[i] = -1

		//向左走 找到第一个没有走过的点
		left := i - 1
		leftNum := 0
		for left > -1 && nums[left] == -1 {
			left--
		}

		//left走到-1，说明左边都走过了，设置成1不影响乘积
		//右边同理
		if left < 0 {
			leftNum = 1
		} else {
			leftNum = nums[left]
		}

		right := i + 1
		rightNum := 0
		for right < length && nums[right] == -1 {
			right++
		}
		if right > length-1 {
			rightNum = 1
		} else {
			rightNum = nums[right]
		}

		// 该值，左，右相乘
		tempProduct := temp * leftNum * rightNum
		dfs(nums, s+1, length, beforeCoins+tempProduct)

		//回溯
		nums[i] = temp
	}
}

//分治法 先算两边再戳破K
//def( i, j ) = max { def( i , k ) + def( k , j )+nums[ i ][ j ][ k ] }

func maxCoins2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 在原来数组的基础上两边多1，代表边界

	newNums := []int{1}
	newNums = append(newNums, nums...)
	newNums = append(newNums, 1)
	length := len(newNums)
	cache := [][]int{}
	for i := 0; i < length; i++ {
		tmp := make([]int, length)
		cache = append(cache, tmp)
	}
	return div(newNums, cache, length, 0, length-1)

}

func div(nums []int, cache [][]int, length, begin, end int) int {
	//开始与结束之间没有气球
	if begin == end-1 {
		return 0
	}
	if cache[begin][end] != 0 {
		return cache[begin][end]
	}
	maxNum := 0
	for i := begin + 1; i < end; i++ {
		temp := div(nums, cache, length, begin, i) + div(nums, cache, length, i, end) + nums[begin]*nums[i]*nums[end]
		if temp > maxNum {
			maxNum = temp
		}
	}
	cache[begin][end] = maxNum
	return maxNum

}

// dp[i][j] 戳破i和j之间气球 乘积的最大值
func maxCoins3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 因为我们只戳破i-j的，所以两边的边界加进来，赋值1不影响乘积
	newNums := []int{1}
	newNums = append(newNums, nums...)
	newNums = append(newNums, 1)
	length := len(newNums)
	dp := [][]int{}
	for i := 0; i < length; i++ {
		tmp := make([]int, length)
		dp = append(dp, tmp)
	}
	//i左j右
	for i := length - 2; i >= 0; i-- {
		for j := i + 2; j < length; j++ {
			maxNum := 0
			//if k=i+1=j  	i，j相邻 maxNum=0
			for k := i + 1; k < j; k++ {
				temp := dp[i][k] + dp[k][j] + newNums[i]*newNums[k]*newNums[j]
				maxNum = max(maxNum, temp)
			}
			dp[i][j] = maxNum
		}
	}
	return dp[0][length-1]
}
