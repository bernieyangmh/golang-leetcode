// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 示例:
// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
// 给定 word = "ABCCED", 返回 true.
// 给定 word = "SEE", 返回 true.
// 给定 word = "ABCB", 返回 false.
package main

import "fmt"

func main() {
	var board [][]byte
	// ['A','B','C','E'],
	// ['S','F','C','S'],
	// ['A','D','E','E']
	tmp := []byte{}
	for _, v := range "ABCE" {
		tmp = append(tmp, byte(v))
	}
	tmp1 := []byte{}
	for _, v := range "SFCS" {
		tmp1 = append(tmp1, byte(v))
	}
	tmp2 := []byte{}
	for _, v := range "ADEE" {
		tmp2 = append(tmp2, byte(v))
	}
	board = append(board, tmp)
	board = append(board, tmp1)
	board = append(board, tmp2)

	fmt.Println(board)
	println(exist(board, "ABCCED"))

}

func exist(board [][]byte, word string) bool {
	// m行
	m := len(board)
	if m == 0 {
		return false
	}
	// n列
	n := len(board[0])
	// 标记已走过的位置
	mark := [][]bool{}
	// 向左,向右,向上,向下
	direction := [][]int{[]int{-1, 0}, []int{0, -1}, []int{0, 1}, []int{1, 0}}

	// 初始化mark
	for i := 0; i < m; i++ {
		oneColunm := []bool(nil)
		for j := 0; j < n; j++ {
			oneColunm = append(oneColunm, false)
		}
		mark = append(mark, oneColunm)
	}
	// 对每个位置,开始深度遍历
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0, word, board, mark, direction) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j, s int, word string, board [][]byte, mark [][]bool, direction [][]int) bool {
	// 递归终止条件
	// 最长路径一定小于等于单词长, 走到单词最后判断结果,返回
	if s == len(word)-1 {
		return board[i][j] == word[s]
	}

	// 和单词不同,直接false不用走了
	if board[i][j] == word[s] {
		// 占住i,j这个位置,搜索不到结果释放掉
		mark[i][j] = true
		// 走四个方向
		for k := 0; k < 4; k++ {
			newx := i + direction[k][0]
			newy := j + direction[k][1]
			// 移动的位置要在board内且没有走过
			if inArea(newx, newy, len(board), len(board[0])) && !mark[newx][newy] {
				if dfs(newx, newy, s+1, word, board, mark, direction) {
					return true
				}
				// 如果写成下面代码,就变成在0,0开始搜索,而不是任意位置开始
				// return dfs(newx, newy, s+1, word, board, mark, direction)
			}
		}
		// 这条路径走不通,回溯, 释放沿途的位置
		mark[i][j] = false
	}
	return false
}

func inArea(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n

}
