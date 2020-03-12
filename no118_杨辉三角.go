package main

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//输入: 5
//输出:
//[
//[1],
//[1,1],
//[1,2,1],
//[1,3,3,1],
//[1,4,6,4,1]
//]


func generate(numRows int) [][]int {
	dp := [][]int{}
	if numRows < 1 {
		return dp
	}
	dp = append(dp, []int{1})

	line := 1
	for line < numRows {

		tmp := make([]int, len(dp[line-1]))
		tmp[0]= 1
		for i:= 1; i < len(dp[line-1]);i++{
			tmp[i] = dp[line-1][i-1] + dp[line-1][i]
		}
		tmp = append(tmp, 1)
		dp = append(dp, tmp)

		line++
	}
	return dp

}


