package main

//给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
//形式上，如果可以找出索引 i+1 < j 且满足 
//(A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。

//暴力算出每个部分的和
func canThreePartsEqualSum(A []int) bool {
	length := len(A)
	dp := make([]int, length)

	dp[0] = A[0]
	for i := 1; i < length; i++ {
		dp[i] = dp[i-1] + A[i]
	}
	// fmt.Println(dp)
	var (
		p1, p2, p3 int
	)

	for left := 1; left < length-1; left++ {
		for right := left; right < length-1; right++ {
			p1 = dp[left-1]
			p2 = dp[right] - dp[left-1]
			p3 = dp[length-1] - dp[right]
			// fmt.Println(left, right, A[:left], p1, A[left:right+1], p2, A[right+1:], p3)
			if p1 == p2 && p2 == p3 {
				return true
			}
		}
	}
	return false
}

func canThreePartsEqualSum2(A []int) bool {

	sum, avg, remainder, part, cnt := 0, 0, 0, 0, 0
	for _, n := range A {
		sum += n
	}
	avg = sum / 3
	remainder = sum % 3

	//有余数的不可能满足条件
	if remainder != 0 {
		return false
	}

	for _, n := range A {
		part += n
		if part == avg {
			cnt++
			part = 0
		}
	}
	return cnt >= 3

}
