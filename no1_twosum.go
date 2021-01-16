package main

import (
	"fmt"
)

// 两数之和

func twoSum(nums []int, target int) []int {
	mapint := make(map[int]int) // index is need completion value, value is index of v
	for i, v := range nums {
		if _, ok := mapint[v]; ok {
			return []int{i, mapint[v]}
		}
		mapint[target-v] = i
	}
	return []int{}
}

func main() {
	nums := []int{3,2,4}
	target := 6
	fmt.Println(twoSum(nums, target))




}
