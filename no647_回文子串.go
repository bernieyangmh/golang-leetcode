package main


// 动态规划，中心扩散 同No.5

// 每次dp=true 计数即可
func countSubstrings(s string) int {
	res := 0
	size := len(s)
	if size <= 1 {
		return size
	}
	var dp [][]bool

	for i := 0; i < size; i++ {
		tmp := []bool{}
		for j := 0; j < size; j++ {
			if j == i {
				tmp = append(tmp, true)
				res++
			} else {
				tmp = append(tmp, false)
			}
		}
		dp = append(dp, tmp)
	}
	for j := 1; j < size;j++{
		for i:=0;i<j;i++{
			if s[i]==s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else{
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if  dp[i][j] {
				res++
			}
		}
	}

	return res
}


func countSubstrings2(s string) int {
	res := 0
	for i:=0; i<len(s);i++{
		res += centernSpread(s, i, i)
		res += centernSpread(s, i,i+1)
	}
	return res
}


func centernSpread(s string, start, end int) int {
	count := 0
	for start>=0 && end < len(s){
		if s[start] != s[end] {
			break
		}
		start--
		end++
		count++
	}
	return count
}

