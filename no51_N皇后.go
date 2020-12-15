package main

import (
	"fmt"
)

//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

func main() {
	fmt.Println(solveNQueens(8))

}

//回溯，减枝
//遍历每一行，看列，主对角线，副对角线满不满足，满足就到下一行，直到能走到最后，把整个结果放进res里
//所有的主对角线有 行号 - 列号 = 固定常数，对于所有的次对角线有 行号 + 列号 = 固定常数
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	// 列是否有
	columns := make([]bool, n) // columns   |
	//主对角线
	d1 := make([]bool, 2*n) // diagonals \
	//副对角线
	d2 := make([]bool, 2*n) // diagonals /

	// 构建一个空的棋盘
	candidate := make([][]rune, n)
	for i := range candidate {
		candidate[i] = make([]rune, n)
		for j := range candidate[i] {
			candidate[i][j] = '.'
		}
	}

	backtracking(candidate, 0, n, columns, d1, d2, &result)
	return result
}

func backtracking(candidate [][]rune, row, n int, columns, d1, d2 []bool, res *[][]string) {
	if row == n {
		var arr []string
		for _, v := range candidate {
			arr = append(arr, string(v))
		}
		*res = append(*res, arr)
		return
	}
	for i := 0; i < n; i++ {
		idx1 := n - i + row
		idx2 := row + i
		if d1[idx1] || d2[idx2] || columns[i] {
			continue
		}
		d1[idx1], d2[idx2], columns[i] = true, true, true
		candidate[row][i] = 'Q'
		backtracking(candidate, row+1, n, columns, d1, d2, res)
		d1[idx1], d2[idx2], columns[i] = false, false, false
		candidate[row][i] = '.'
	}
}
