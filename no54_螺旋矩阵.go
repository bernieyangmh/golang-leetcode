package main


// 题简单，要细心。。。
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// 初始化
	passed := [][]bool{}
	for i:=0;i<len(matrix);i++{
		tmp := make([]bool, len(matrix[0]))
		passed = append(passed, tmp)
	}
	// 开始的位置和方向
	i:=0
	j:=0
	dir := 1

	// 把第一步放进结果和已走过路径中
	res := []int{matrix[i][j]}
	passed[i][j] = true


	for {
		//根据当前位置和方向判断下一步的位置
		nextI, nextJ := walk(i, j, dir)
		// 如果可以,进入下一步;如果遇到边界或者已经走过,改变方向
		if check(matrix, passed, nextI, nextJ) {
			res = append(res, matrix[nextI][nextJ])
			passed[nextI][nextJ] = true
			i, j = nextI, nextJ
			continue
		} else {
			// 换方向走
			dir = nextDir(dir)
			nextI, nextJ = walk(i, j, dir)

			if check(matrix, passed, nextI, nextJ) {
				res = append(res, matrix[nextI][nextJ])
				passed[nextI][nextJ] = true
				i, j = nextI, nextJ
				continue
			} else {
				// 换方向还是不行, 说明走到头了
				return res
			}
		}
	}
	return res
}

// 边界检查
func check(matrix [][]int, passed [][]bool, i, j int) bool {
	// 左、上边界
	if i < 0 || j < 0{
		return false
	}
	// 右、下边界
	if i >= len(matrix) || j >= len(matrix[0]) {
		return false
	}
	// 已走过
	if passed[i][j] {
		return false
	}
	return true
}

// 方向变更
// 1右,2下,3左,4上
func nextDir(dir int)int{
	switch dir {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 4
	case 4:
		return 1
	}
	return 0
}

// 下一步的位置
func walk(i, j, dir int) (int, int) {
	switch dir {
	case 1:
		return i, j + 1
	case 2:
		return i + 1, j
	case 3:
		return i, j - 1
	case 4:
		return i -1 , j
	}
	return 0, 0
}