package main

import (
	"fmt"
)

func main() {
	arr := []int{1,1,2,2,3,3,4,4,5,5,123,123,421,432,421}
	fmt.Println(findSingle(arr))
}

func findSingle(arr []int) int {
	res := 0
	for i:=0; i<len(arr);i++{
		res ^= arr[i]
	}
	return res
}