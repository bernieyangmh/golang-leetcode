package main

import (
	"fmt"
)

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
	nums := []int{2, 7, 4, 8}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(5/2)

}
