package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 1}
	k := 2
	fmt.Println(GetLeastNumbers_Solution(a, k))
	//[0,0,1,1,2,2,2,3]
}

func GetLeastNumbers_Solution(input []int, k int) []int {
	if k == 0 || k > len(input) {

		return []int{}
	}
	qs(input, 0, len(input)-1)
	return input[:k]
}

func qs(arr []int, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		qs(arr, l, p-1)
		qs(arr, p+1, r)
	}
}

func partition(arr []int, l, r int) int {
	p := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] <= p {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
