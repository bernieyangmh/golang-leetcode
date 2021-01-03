package main

import (
	"fmt"
)

func main()  {
	a := make(map[*int]int)
	b := &[]int{1}[0]
	a[b]=1

	a[b]=2
	fmt.Println(a)
}