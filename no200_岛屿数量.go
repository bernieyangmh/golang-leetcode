// 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

// 示例 1:
// 11110
// 11010
// 11000
// 00000
// 输出: 1

// 示例 2:
// 11000
// 11000
// 00100
// 00011
// 输出: 3

package main

import (
	"fmt"
)

// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]

func main() {
	// a := "10"
	// for _, value := range a {
	// 	fmt.Println(value)
	// }

	a := [][]byte{}
	a = append(a, []byte{49, 49, 49, 49, 48})
	a = append(a, []byte{49, 49, 48, 49, 48})
	a = append(a, []byte{49, 49, 48, 48, 48})
	a = append(a, []byte{48, 48, 48, 48, 48})
	// a = append(a, []byte{48, 48, 48, 48, 48})
	fmt.Println(numIslands(a))

}

// 遍历grid, 当遇到陆地,从该点蔓延将所有遇到的陆地都变成水,再继续遍历,再遇到的陆地一定和之前的陆地没有连接,是新的岛屿
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var cnt int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 49 {
				cnt++
				grid = fillWater(grid, i, j, 0)
			}
		}
	}
	return cnt
}

// 1左 2右 3上 4下
func fillWater(grid [][]byte, i, j, from int) [][]byte {
	// fmt.Println(grid, i, j, from)
	grid[i][j] = 48
	// 不是从左边来的,判断左边
	if from != 1 && i != 0 && grid[i-1][j] == 49 {
		grid = fillWater(grid, i-1, j, 2)
	}
	//
	if from != 2 && i != len(grid)-1 && grid[i+1][j] == 49 {
		grid = fillWater(grid, i+1, j, 1)
	}

	if from != 3 && j != 0 && grid[i][j-1] == 49 {
		grid = fillWater(grid, i, j-1, 4)
	}
	if from != 4 && j != len(grid[0])-1 && grid[i][j+1] == 49 {
		grid = fillWater(grid, i, j+1, 3)
	}
	return grid
}

// type UnionFind struct {
// 	Count  int
// 	Parent []int
// 	Rank   []int
// }

// func newUnionFind(grid [][]byte) *UnionFind {
// 	count := 0
// 	m := len(grid)
// 	n := len(grid[0])
// 	parent := []int{}
// 	rank := []int{}

// 	for i := 0; i < m*n; i++ {
// 		parent = append(parent, 0)
// 		rank = append(rank, -1)
// 	}

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if grid[i][j] == 49 {
// 				parent[i*n+j] = i*n + j
// 				count++
// 			}
// 			rank[i*n+j] = 0
// 		}
// 	}
// 	Uf := new(UnionFind)
// 	Uf.Count = count
// 	Uf.Parent = parent
// 	Uf.Rank = rank
// 	return Uf
// }

// func (uf *UnionFind) find(i int) int {
// 	if uf.Parent[i] != i {
// 		uf.Parent[i] = uf.find(uf.Parent[i])
// 	}
// 	return uf.Parent[i]

// }

// func (uf *UnionFind) union(x, y int) {
// 	rootX := uf.find(x)
// 	rootY := uf.find(y)
// 	if rootX != rootY {
// 		if uf.Rank[rootX] > uf.Rank[rootY] {
// 			uf.Parent[rootY] = rootX
// 		} else if uf.Rank[rootX] < uf.Rank[rootY] {
// 			uf.Parent[rootX] = rootY
// 		} else {
// 			uf.Parent[rootY] = rootX
// 			uf.Rank[rootX] += 1
// 		}
// 		uf.Count -= 1
// 	}
// }

// func numIslands2(grid [][]byte) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}
// 	nr := len(grid)
// 	nc := len(grid[0])
// 	uf := newUnionFind(grid)
// 	for r := 0; r < nr; r++ {
// 		for c := 0; c < nc; c++ {
// 			if grid[r][c] == 49 {
// 				grid[r][c] = 48
// 				if r-1 >= 0 && grid[r-1][c] == 49 {
// 					uf.union(r*nc+c, (r-1)*nc+c)
// 				}
// 				if r+1 < nr && grid[r+1][c] == 49 {
// 					uf.union(r*nc+c, (r+1)*nc+c)
// 				}
// 				if c-1 >= 0 && grid[r][c-1] == 49 {
// 					uf.union(r*nc+c, r*nc+c-1)
// 				}
// 				if c+1 < 0 && grid[r][c+1] == 49 {
// 					uf.union(r*nc+c, r*nc+c+1)
// 				}
// 			}
// 		}
// 	}
// 	return uf.Count
// }
