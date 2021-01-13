package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i >= v {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Partition(input []int, l, r int) int {
	pivot := input[r-1]
	i := l

	for j := l; j < r-1; j++ {
		if input[j] < pivot {

			input[i+1], input[j] = input[j], input[i+1]
		}
	}
	input[i], input[r-1] = input[r-1], input[i]
	return i
}

func main() {
	//[4,5,1,6,2,7,3,8],8
	a := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers_Solution(a, 8))

}

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	ret := make([]int, k)
	if k == 0 || k > len(input) {
		return ret
	}
	l, r := 0, len(input)-1
	for l <= r {
		p := partition(input, l, r)
		fmt.Println(p)
		if p+1 == k {
			break
		}
		if p+1 < k {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	for i := 0; i < k; i++ {
		ret[i] = input[i]
	}

	return ret
}

func partition(input []int, l, r int) int {
	pNum := input[r]
	s := l - 1

	for l < r {
		if input[l] < pNum {
			s++
			swap(input, l, s)
		}
		l++
	}
	s++
	swap(input, s, r)
	return s
}

func swap(arr []int, n1, n2 int) {
	arr[n1], arr[n2] = arr[n2], arr[n1]

}
