package collections

//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

func jumpFloor(number int) int {
	// write code here
	return Fibonacci(number)

}

func Fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func Fibonacci3(n int) int {
	dp_res := n
	dp_0, dp_1 := 1, 1
	for i := 2; i <= n; i++ {
		dp_res = dp_0 + dp_1
		dp_0, dp_1 = dp_1, dp_res
	}
	return dp_res
}
