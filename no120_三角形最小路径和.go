package main

//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//例如，给定三角形：
//[
//[2],
//[3,4],
//[6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := []int{triangle[0][0]}
	for i := 1; i < len(triangle); i++ {
		// 每一行的头的路径和是该首位数+之前结果的首位数
		tmp := make([]int, len(triangle[i]))
		tmp[0] = triangle[i][0] + dp[0]
		for j := 1; j < len(triangle[i])-1; j++ {
			//中间数的路径和是左右结果最小值+该点数值
			tmp[j] = triangle[i][j] + min(dp[j], dp[j-1])
		}
		//尾数和是尾数+之前结果的尾数
		tmp[len(tmp)-1] = triangle[i][len(triangle[i])-1] + dp[len(dp)-1]
		dp = tmp
	}
	minPath := 1<<31 - 1
	for _, i := range dp {
		minPath = min(minPath, i)
	}
	return minPath
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
