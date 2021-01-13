package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2}
	arr2 := []int{1, 0}

	merge(arr1, len(arr1), arr2, len(arr2))
}

func merge(A []int, m int, B []int, n int) {
	// write code here

	if m == 0 && n == 0 {
		return
	}

	p1, p2 := m-1, n-1
	cur := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if A[p1] > B[p2] {
			A[cur] = A[p1]
			p1--
		} else {
			A[cur] = B[p2]
			p2--
		}
		cur--
	}

	for p2 >= 0 {
		A[cur] = B[p2]
		cur--
		p2--
	}

	return
}
