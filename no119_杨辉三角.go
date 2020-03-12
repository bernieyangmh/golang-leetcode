package main


func getRow(rowIndex int) []int {
	dp := []int{}
	if rowIndex < 0 {
		return dp
	}
	dp = append(dp, 1)

	line := 1

	for line <= rowIndex {
		tmp := []int{1}
		for i:=1;i<len(dp);i++{
			tmp =append(tmp, dp[i-1]+dp[i])
		}
		tmp = append(tmp, 1)
		dp = tmp
		line++
	}

	return dp

}

func getRow2(rowIndex int) []int {
	ans := []int{1}
	n := rowIndex
	pre := 1

	for k:=1;k<=n;k++{
		cur := pre * (n-k+1)/k
		ans = append(ans, cur)
		pre = cur
	}
	return ans
}
