// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 示例:
// 现有矩阵 matrix 如下：
// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。
// 给定 target = 20，返回 false。
package main

import "fmt"

func main() {
	a := []int{1, 4, 7, 11, 15}

	// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]

	fmt.Println(binarySearch(a, 5, 0, 4))
}

//掐头去尾,然后二叉搜索
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for _, m := range matrix {
		if m[0] == target {
			return true
		}
		if m[len(m)-1] == target {
			return true
		}
		if m[0] < target && m[len(m)-1] < target {
			if binarySearch(m, target, 0, len(m)) {
				return true
			}
		}
	}
	return false
}

func binarySearch(arr []int, target, s, e int) bool {
	if arr[(s+e)/2] == target {
		return true
	}
	if s >= e {
		return false
	}
	if arr[(s+e)/2] <= target {
		return binarySearch(arr, target, (s+e)/2+1, e)

	} else {
		return binarySearch(arr, target, s, (s+e)/2-1)
	}

}

// 左上小右下大, 选定一个初始点,然后选择一个方向前进
// 选左上角，往右走和往下走都增大，不能选
// 选右下角，往上走和往左走都减小，不能选
// 选左下角，往右走增大，往上走减小，可选
// 选右上角，往下走增大，往左走减小，可选

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	point_line := m - 1
	point_column := 0

	for point_line >= 0 && point_column < n {
		if matrix[point_line][point_column] > target {
			point_line--
		} else if matrix[point_line][point_column] < target {
			point_column++
		} else {
			return true
		}
	}
	return false

}
