package main

import (
	"fmt"
)

func main()  {
	a := []int{3,2,1}
	k := 2
	fmt.Println(getLeastNumbers(a, k))
	//[0,0,1,1,2,2,2,3]
}

func getLeastNumbers(arr []int, k int) []int {
	stack := []int{}
	if k == 0 {
		return stack
	}
	for _, n := range arr{
		// 如果比栈顶大,跳过
		if len(stack) > 0 && n >= stack[len(stack)-1] {
			if len(stack) < k {
				stack = append(stack, n)
			}
			// 栈没有元素或比栈顶小
		} else if len(stack) == 0 {
			stack = append(stack, n)
		} else if n <= stack[len(stack)-1] {
			// 找到最小的
			for i:= len(stack)-1;i >= 0; i--{
				if stack[i] <= n {
					p2 := append([]int{}, stack[i+1:]...)
					fmt.Println(stack, p2, n)
					stack = append(stack[:i+1], n)
					fmt.Println(stack)
					stack = append(stack, p2...)
					fmt.Println(stack)
					fmt.Println()
					break
				}
				if n < stack[0] {
					stack = append([]int{n}, stack...)
				}
			}
			if len(stack) > k {
				stack = stack[:k]
			}
		}
	}
	return stack
}