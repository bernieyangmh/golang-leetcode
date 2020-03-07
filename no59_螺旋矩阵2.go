package main

import (
	"fmt"
)

//给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//示例:
//输入: 3
//输出:
//[
//[ 1, 2, 3 ],
//[ 8, 9, 4 ],
//[ 7, 6, 5 ]
//]

func main() {
	fmt.Println(generateMatrix(3))

}

// 1.每走n-i+1圈改变方向,即4,3,2,1
// 规律
// n=1 => 1 1 1 0
// n=2 => 2,2,2,1,1,0
// n=3 => 3,3,3,2,2,1,1,0
// n=4 => 4,4,4,3,3,2,2,1,1,0
// n=5 => 5,5,5,4,4,3,3,2,2,1,1,0
// 第一次变3次方向之后步数-1, 其他都是变两个方向就-1
// 去掉一开始的一条边, 每走两条边,长宽已走过的距离就多1, 收缩一圈
// 细节较多
func generateMatrix(n int) [][]int {
	// 位置,方向,数字
	i, j, dir, num := 0, 0, 1, 1
	// 用来计算剩余可以走的步数，一开始要在正方形的边唱之外
	leftB, rightB, topB, bottomB := -1, n, -1, n
	square := [][]int{}
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		square = append(square, tmp)
	}

	// 第一次向右走
	for j < n {
		square[i][j] = num
		j++
		num++
	}
	j--
	// 下一个位置的判断最重要
	for num <= n*n {
		nextI, nextJ := walk(i, j, dir)
		//fmt.Println(i, j,nextI,nextJ, dir)
		//如果沿着当前方向去下一个位置碰到边界，就换方向
		if check(nextI, nextJ, dir, leftB, rightB, topB, bottomB) {
			i, j = nextI, nextJ
			square[i][j] = num
		} else {
			// 上走完，上边收缩，右走完，右边收缩...
			switch dir {
			case 1:
				topB++
			case 2:
				rightB--
			case 3:
				bottomB--
			case 4:
				leftB++
			}
			dir = nextDir(dir)
			i, j = walk(i, j, dir)
			square[i][j] = num

		}
		num++
	}
	return square
}

func check(i, j, dir, leftB, rightB, topB, bottomB int) bool {
	switch dir {
	case 1:
		return j < rightB
	case 2:
		return i < bottomB
	case 3:
		return j > leftB
	case 4:
		return i > topB
	}
	return false
}

// 1右,2下,3左,4上
func nextDir(dir int) int {
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
		return i - 1, j
	}
	return 0, 0
}
