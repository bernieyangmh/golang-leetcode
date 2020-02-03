// 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 示例:
// X X X X
// X O O X
// X X O X
// X O X X
// 运行你的函数后，矩阵变为：
// X X X X
// X X X X
// X X X X
// X O X X
// 解释:
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

package main

func main() {
	a := "XO"
	for _, v := range a {
		println(v)
	}

}

// 暴力判断,会超时
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 79 {
				isClose(board, i, j, 0)
			}
		}
	}
}

// 判断o是否被x包围,i,j的位置是o, mark 代表我们已经判断过的O的位置
func isClose(board [][]byte, i, j, from int) bool {
	//from 代表从哪个方向移动过来的,我们不需要再回头判断之前的点是不是符合
	flag := false
	// 如果当前位置在左边界
	if from != 1 {
		if j == 0 {
			return false
		}
		// 不在的话左边是不是o,是判断左边的o行不行
		if board[i][j-1] == 79 {
			flag = isClose(board, i, j-1, 2)
			if !flag {
				return flag
			}
		}
	}

	if from != 2 {
		// 如果当前位置在右边界
		if j == len(board[0])-1 {
			return false
		}
		if board[i][j+1] == 79 {
			flag = isClose(board, i, j+1, 1)
			if !flag {
				return flag
			}
		}
	}

	if from != 3 {
		// 如果当前位置在上边界
		if i == 0 {
			return false
		}
		if board[i-1][j] == 79 {
			flag = isClose(board, i-1, j, 4)
			if !flag {
				return flag
			}
		}
	}

	if from != 4 {
		// 如果当前位置在下边界
		if i == len(board)-1 {
			return false
		}
		if board[i+1][j] == 79 {
			flag = isClose(board, i+1, j, 3)
			if !flag {
				return flag
			}
		}
	}
	// 走到这里,说明该位置及其相邻的o被x包围
	board[i][j] = 88
	return true
}

// 从四条边开始走,遇到O换成Y, 不和边界联通的O仍为O
// 结束之后遍历每个格子,仍为O的换成X, Y的换成O
func solve2(board [][]byte) {
	if len(board) == 0 {
		return
	}
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 79 {
			borderO(board, 0, j, 0)
		}
		if board[len(board)-1][j] == 79 {
			borderO(board, len(board)-1, j, 0)
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][0] == 79 {
			borderO(board, i, 0, 0)
		}
		if board[i][len(board[0])-1] == 79 {
			borderO(board, i, len(board[0])-1, 0)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 79 {
				board[i][j] = 88
			}
			if board[i][j] == 89 {
				board[i][j] = 79
			}
		}
	}
}

// 把边界联通的O都换成Y
func borderO(board [][]byte, i, j, from int) {
	// 换字母
	board[i][j] = 89
	// 如果不是从左边来
	if from != 1 {
		// 边界上的就不用重复判断了
		if j-1 > 0 && board[i][j-1] == 79 {
			borderO(board, i, j-1, 2)
		}
	}

	if from != 2 {
		if j+1 < len(board[0])-1 && board[i][j+1] == 79 {
			borderO(board, i, j+1, 1)
		}
	}

	if from != 3 {
		if i-1 > 0 && board[i-1][j] == 79 {
			borderO(board, i-1, j, 4)
		}
	}

	if from != 4 {
		if i+1 < len(board)-1 && board[i+1][j] == 79 {
			borderO(board, i+1, j, 3)
		}
	}
}
