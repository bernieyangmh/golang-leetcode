package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 6, 21, 3, 4, 6, 8, 9, 3, 1, 3, 5, 6, 7, 3, 1, 312, 4, 5345, 123, 123}

	foo2(arr)
	fmt.Println(arr)
}


func BubbleAsort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func foo(values []int) {
	values[3] = 33
}

func foo2(values []int) {
	values = append(values, 1)
	values[3] = 33
}