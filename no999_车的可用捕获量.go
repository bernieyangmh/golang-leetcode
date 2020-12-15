package main

func numRookCaptures(board [][]byte) int {
	dir := [][]int{}
	dir = append(dir, []int{-1, 0})
	dir = append(dir, []int{1, 0})
	dir = append(dir, []int{0, 1})
	dir = append(dir, []int{0, -1})

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				res := 0
				for _, d := range dir {
					if run(board, i, j, d) {
						res++
					}
				}
				return res
			}

		}
	}
	return 0
}

func run(board [][]byte, x, y int, dir []int) bool {
	i := x
	j := y
	for inArea(i, j) {
		if board[i][j] == 'B' {
			break
		}
		if board[i][j] == 'p' {
			return true
		}
		i += dir[0]
		j += dir[1]
	}
	return false

}

func inArea(i, j int) bool {
	return i >= 0 && i < 8 && j >= 0 && j < 8

}
