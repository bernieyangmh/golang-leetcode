package main

import (
	"fmt"
)

var res int

func main()  {
	fmt.Println(totalNQueens(8))
}

func totalNQueens(n int) int {
	res = 0
	col := make([]bool, n)
	d1 := make([]bool, 2*n)
	d2 := make([]bool, 2*n)

	backtracking(0, n, col, d1, d2)
	return res

}

func backtracking(row, n int, col, d1, d2 []bool) {
	if row == n {
		res++
		return
	}
	for i := 0; i < n; i++ {
		index1 := n - i + row
		index2 := i + row
		if d1[index1] || d2[index2] || col[i] {
			continue
		}
		d1[index1], d2[index2], col [i] = true, true, true
		backtracking(row+1, n, col, d1, d2)
		d1[index1], d2[index2], col [i] = false, false, false

	}
}
