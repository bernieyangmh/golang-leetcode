package main

import (
	"fmt"
	"math/rand"
)

// d快排

func main() {
	arr := []int{1, 4, 6, 21, 3, 4, 6, 8, 9, 3, 1, 3, 5, 6, 7, 3, 1, 312, 4, 5345, 123, 123}

	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
