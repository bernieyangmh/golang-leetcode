// 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
// 示例:
// 输入:
// [
//   ["1","0","1","0","0"],
//   ["1","0","1","1","1"],
//   ["1","1","1","1","1"],
//   ["1","0","0","1","0"]
// ]
// 输出: 6
package main

import "fmt"

func main() {
	a := [][]byte{}
	a = append(a, []byte{48, 48, 48, 48, 48})
	a = append(a, []byte{48, 49, 49, 49, 48})
	a = append(a, []byte{48, 49, 49, 49, 48})
	a = append(a, []byte{49, 49, 49, 49, 48})
	fmt.Println(maximalRectangle3(a))

}

// 运用dp 找到该行每个位置前连续1的长度为宽
// 每个位置向上, 找到以该位置为右下角的每个矩阵的面积
// area = min(width) * height
// 宽的最小值为向上每个点连续宽的最小值, 高是该点向上移动的距离
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	dp := [][]int{}
	m := len(matrix)
	n := len(matrix[0])
	maxAera := 0

	for i := 0; i < m; i++ {
		// 因为我们每次只需要上面行和当前行的数据,dp的生成可以直接放到遍历里
		dp = append(dp, make([]int, n))
		for j := 0; j < n; j++ {
			// 0够不成,不考虑
			if string(matrix[i][j]) == "1" {
				if j > 0 {
					dp[i][j] = dp[i][j-1] + 1
				} else {
					dp[i][j] = 1
				}
				width := dp[i][j]
				for k := i; k >= 0; k-- {
					width = min(width, dp[k][j])
					maxAera = max(maxAera, width*(i-k+1))
				}
			}
		}
	}

	return maxAera

}

func maximalRectangle2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])
	maxarea := 0
	dp := make([]int, n)

	// 将每一行转化为类似84的柱状图,然后求面积
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if string(matrix[i][j]) == "1" {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		maxarea = max(maxarea, largestRectangleArea4(dp))
	}
	return maxarea
}

func largestRectangleArea4(heights []int) int {
	stack := []int{}
	dummyH := []int{0}
	dummyH = append(dummyH, heights...)
	dummyH = append(dummyH, 0)
	maxAera := 0

	for i := 0; i < len(dummyH); i++ {
		for len(stack) != 0 && dummyH[stack[len(stack)-1]] > dummyH[i] {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxAera = max(maxAera, (i-stack[len(stack)-1]-1)*dummyH[tmp])
		}
		stack = append(stack, i)
	}
	return maxAera

}

func maximalRectangle3(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	maxArea := 0

	left := make([]int, n)
	right := make([]int, n)
	height := make([]int, n)

	// 每一个位置的右边最初都在最右端
	for i, _ := range right {
		right[i] = n
	}

	for i := 0; i < m; i++ {
		cur_left, cur_right := 0, n
		for j := 0; j < n; j++ {
			// j这个点在当前i行上面有多少个连续的1, 高度
			if string(matrix[i][j]) == "1" {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// j这个点往左走第一个0的坐标, 没有-1代表0右边1的坐标,避免-1
		// 方便计算,下面计算面积前距离也没+1
		for j := 0; j < n; j++ {
			if string(matrix[i][j]) == "1" {
				left[j] = max(left[j], cur_left)
			} else {
				left[j], cur_left = 0, j+1
			}
		}

		// j这一列最右边0的坐标
		for j := n - 1; j >= 0; j-- {
			if string(matrix[i][j]) == "1" {
				right[j] = min(cur_right, right[j])
			} else {
				right[j], cur_right = n, j
			}
		}
		for j := 0; j < n; j++ {
			maxArea = max(maxArea, (right[j]-left[j])*height[j])
		}
		fmt.Println(height, left, right)
	}
	return maxArea
}


// 边界数值更易懂
func maximalRectangle4(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	maxArea := 0

	left := make([]int, n)
	right := make([]int, n)
	height := make([]int, n)

	for i:=0;i<n;i++{
		left[i] = -1
	}
	for i:=0;i<n;i++{
		right[i] = n
	}

	for r:=0; r<m;r++{
		for c :=0;c<n;c++{
			if string(matrix[r][c]) == "1" {
				height[c]++
			}else {
				height[c] = 0
			}
		}
		cur_left := -1
		for c:=0;c<n;c++{
			if string(matrix[r][c]) == "1" {
				// 从j往上这一列,最左边的边的位置取决于 j左边 最右边没有0的列
				left[c] = max(left[c], cur_left)
			}else {
				left[c], cur_left = -1, c
			} 
		}

		cur_right := n
		for c:=n-1;c>=0;c--{
			if string(matrix[r][c]) == "1" {
				// 从j往上这一列,最右边的边的位置取决于 j右边 最左没有0的列
				right[c] = min(right[c], cur_right)
			}else {
				right[c], cur_right = n, c
			} 
		}
		for c:=n-1;c>=0;c--{
			maxArea = max(maxArea, (right[c]-left[c]-1) * height[c] )
		}
	}
	return maxArea
}












func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
