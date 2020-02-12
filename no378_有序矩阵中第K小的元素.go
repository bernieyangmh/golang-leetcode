// 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
// 请注意，它是排序后的第k小元素，而不是第k个元素。
// 示例:
// matrix = [
//    [ 1,  5,  9],
//    [10, 11, 13],
//    [12, 13, 15]
// ],
// k = 8,
// 返回 13。
// 说明:
// 你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n2 。

package main

func kthSmallest(matrix [][]int, k int) int {
	row := len(matrix)
	col := len(matrix[0])
	left := matrix[0][0]
	right := matrix[row-1][col-1]

	for left < right {
		mid := (left + right) / 2
		count := findSmallCnt(matrix, mid, row, col)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return right
}

func findSmallCnt(matrix [][]int, mid, row, col int) int {
	i := row - 1
	j := 0
	count := 0
	for i >= 0 && j < col {
		// 向右移动,加一整行
		if matrix[i][j] <= mid {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count

}
