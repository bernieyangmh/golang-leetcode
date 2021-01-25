package main

// 第K大的数

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(FindKthMax(arr, 4))
}

func FindKthMax(arr []int, k int) int {
	for i := 1; i < k; i++ {
		for j := i; j > 0; i-- {
			if arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	for i := k; i < len(arr); i++ {
		if arr[i] > arr[k-1] {
			arr[k-1] = arr[i]
			for j := k - 1; k > 0; k-- {
				if arr[j] > arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
	}
	return arr[k-1]
}
