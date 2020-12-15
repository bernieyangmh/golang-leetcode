package main

import (
	"fmt"
)

func main() {
	a := []int{186, 419, 83, 408}
	fmt.Println(coinChange(a, 1024))

}



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