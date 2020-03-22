// 给你一个 n 行 m 列的二维网格 grid 和一个整数 k。你需要将 grid 迁移 k 次。
// 每次「迁移」操作将会引发下述活动：
// 位于 grid[i][j] 的元素将会移动到 grid[i][j + 1]。
// 位于 grid[i][n - 1] 的元素将会移动到 grid[i + 1][0]。
// 位于 grid[m - 1][ - 1] 的元素将会移动到 grid[0][0]。
// 请你返回 k 次迁移操作后最终得到的 二维网格。

package main

import (
	"fmt"
)

func main() {
	grid := [][]int{}
	grid = append(grid, []int{1})
	// grid = append(grid, []int{4, 5, 6})
	// grid = append(grid, []int{7, 8, 9})
	grid = shiftGrid(grid, 1)
	fmt.Println(grid)
}

// 相当于一维数组移位
func shiftGrid(grid [][]int, k int) [][]int {
	var oneArr []int

	// 依次添加到1维数组
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			oneArr = append(oneArr, grid[i][j])
		}
	}
	// 移动m*n 次相当于没有移动,取 除m*n的 余数
	k = k % (len(grid) * len(grid[0]))


	// 右移k项 经典的三次翻转法，我们可以这么做：
	//							1234567
	// 先把[0, n - k - 1]翻转  	5432167
	// 然后把[n - k, n - 1]翻转	5432176
	// 最后把[0, n - 1]翻转		6712345
	reverse(oneArr, 0, len(oneArr)-k-1)
	reverse(oneArr, len(oneArr)-k, len(oneArr)-1)
	reverse(oneArr, 0, len(oneArr)-1)

	// 构成二维数组
	var line int
	var Column int
	for i := 0; i < len(oneArr); i++ {
		if Column < len(grid[0]) {
			grid[line][Column] = oneArr[i]
			Column += 1
		} else {
			Column = 0
			line += 1
			grid[line][Column] = oneArr[i]
			Column += 1
		}
	}
	return grid
}

// 反转
func reverse(arr []int, start, end int) {
	for start < end {
		tmp := arr[start]
		arr[start] = arr[end]
		arr[end] = tmp
		start += 1
		end -= 1
	}
}
