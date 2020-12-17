package main

//121. Best Time to Buy and Sell Stock
//122. Best Time to Buy and Sell Stock II
//123. Best Time to Buy and Sell Stock III
//188. Best Time to Buy and Sell Stock IV
//309. Best Time to Buy and Sell Stock with Cooldown
//714. Best Time to Buy and Sell Stock with Transaction Fee

//6道股票题统一用相同的动态规划思想解决
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems

//summary

// 1.状态定义
//dp[n][k][0/1] 第n天还能进行k次交易，手中是否持有股票，当天截止的全部收益

// 2.状态转移方程
//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
//i天没有股票，可能是i-1天就没有股票，也可能是i-1天卖掉了股票

//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
//i天持有股票，可能是i-1天就持有股票， 也可能是i-1天买了股票

//从k开始一直算到0， No123两次的时候先算第二天，No188k次的时候先算第k天

// 3.初始化 不可能买股票时的收益是负无穷(天数不存在或无交易)，第一天不买时max一定是0，第一天买时max一定是-prices[0]
//dp[-1][k][0] = 0
//dp[-1][k][1] = 1<<31*-1
//dp[i][0][0] = 0
//do[i][0][1] = 1<<31*-1



//No121 -- 交易只有一次
// 非动态规划做法 找到最小值，当前股价大于最小值->更新最大利润

// 只有一次交易，不需要k
func maxProfit121(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, 1<<31*-1

	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		// 与No.122的区别，就一次买的机会，减股票价格就够了
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

//No.122 -- 交易不限次数
// 非动态规划做法 只有两天有利润cur - prev > 0，就买卖 (1,3,5)这种排序5-3+3-1 == 5-1

// 不限交易次数，k永远是正无穷，可以忽略
func maxProfit122(prices []int) int {
	dp_i_0, dp_i_1 := 0, 1<<31*-1
	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		// 唯一的区别，在卖出股票的收益基础上买入股票
		dp_i_1 = max(dp_i_1, dp_i_0-prices[i])
	}
	return dp_i_0
}

// No.123 -- 交易最多两次
// k是2，4个变量就可表示两天的买卖利润
func maxProfit123(prices []int) int {
	dp_i_1_0, dp_i_1_1, dp_i_2_0, dp_i_2_1 := 0, 1<<31*-1, 0, 1<<31*-1
	for i := 0; i < len(prices); i++ {
		//第二次卖=第二次买+股价
		dp_i_2_0 = max(dp_i_2_0, dp_i_2_1+prices[i])
		//第二次买=第一次卖-股价
		dp_i_2_1 = max(dp_i_2_1, dp_i_1_0-prices[i])
		//第一次卖=第一次买+股价
		dp_i_1_0 = max(dp_i_1_0, dp_i_1_1+prices[i])
		//第一次买=-股价
		dp_i_1_1 = max(dp_i_1_1, -prices[i])
	}
	return dp_i_2_0
}

// No.188 -- 最多交易k次
//交易最多天数的1/2次，//k>= 1/2 ~ no122 交易无限次
func maxProfit188(k int, prices []int) int {
	minNum := 1 << 31 * -1

	if k >= (len(prices) >> 1) {
		dp_i_0, dp_i_1 := 0, minNum
		for i := 0; i < len(prices); i++ {
			dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
			// 唯一的区别，在卖出股票的收益基础上买入股票
			dp_i_1 = max(dp_i_1, dp_i_0-prices[i])
		}
		return dp_i_0
	}
	dp_i_k_0 := make([]int, k+1)
	dp_i_k_1 := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		dp_i_k_1[i] = minNum
	}

	for _, p := range prices {
		for j := k; j > 0; j-- {
			dp_i_k_0[j] = max(dp_i_k_0[j], dp_i_k_1[j]+p)
			dp_i_k_1[j] = max(dp_i_k_1[j], dp_i_k_0[j-1]-p)
		}
	}
	return dp_i_k_0[k]
}

// No.309 卖后有冷却期
// 对于冷却期，多一天保存冷却期前的利润，其他与No.122无限次相同
func maxProfit309(prices []int) int {
	dp_prev_0, dp_i_0, dp_i_1 := 0, 0, 1<<31*-1
	for _, p := range prices {
		dp_i_0_origin := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+p)
		dp_i_1 = max(dp_i_1, dp_prev_0-p)
		dp_prev_0 = dp_i_0_origin
	}
	return dp_i_0
}


// No.714买卖有手续费
// 手续费，在买或卖时扣钱
func maxProfit714(prices []int, fee int) int {
	dp_i_0, dp_i_1 := 0, 1<<31*-1

	for _, p := range prices{
		old0 := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+p)
		dp_i_1 = max(dp_i_1, old0-p-fee)
	}
	return dp_i_0
}

func maxProfit7142(prices []int, fee int) int {
	dp_i_0, dp_i_1 := 0, 1<<31*-1

	for _, p := range prices{
		old0 := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+p-fee)
		dp_i_1 = max(dp_i_1, old0-p)
	}
	return dp_i_0
}


func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
