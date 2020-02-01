package main

import "fmt"

func main() {
	matrix := [][]int{}
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	rotate(matrix)
	fmt.Println(matrix)
}

// 每次旋转一层正方形的位置, 像旋转洋葱的每一层
func rotate(matrix [][]int) {
	// 正方形的边长
	m := len(matrix)
	// 当前正方形离最外层正方形的距离
	n := 0

	for m-n*2 > 0 {
		rotateLayer(matrix, n, m)
		n++
	}
}

// 置换一层正方形的边的位置
func rotateLayer(matrix [][]int, n, m int) {
	// 初始化
	// 左上位置
	leftTop := []int{n, n}
	// 右上位置
	rightTop := []int{n, m - 1 - n}
	// 左下位置
	leftBottom := []int{m - 1 - n, n}
	// 右下位置
	rightBottom := []int{m - 1 - n, m - 1 - n}

	// 当左上角的位置移动到最初的右上角,停止循环,判断其他角的条件也行, 其实循环了m-n-1次
	for leftTop[1] < m-1-n {
		// 左上和右上互换位置
		matrix[leftTop[0]][leftTop[1]], matrix[rightTop[0]][rightTop[1]] = matrix[rightTop[0]][rightTop[1]], matrix[leftTop[0]][leftTop[1]]
		// 左下和右下互换位置
		matrix[leftBottom[0]][leftBottom[1]], matrix[rightBottom[0]][rightBottom[1]] = matrix[rightBottom[0]][rightBottom[1]], matrix[leftBottom[0]][leftBottom[1]]
		// 左上和右下互换位置
		matrix[leftTop[0]][leftTop[1]], matrix[rightBottom[0]][rightBottom[1]] = matrix[rightBottom[0]][rightBottom[1]], matrix[leftTop[0]][leftTop[1]]

		// 左上右移, 右上下移,右下左移, 左下上移
		leftTop[1] += 1
		rightTop[0] += 1
		rightBottom[1] -= 1
		leftBottom[0] -= 1
	}

}
