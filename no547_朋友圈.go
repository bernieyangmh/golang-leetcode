// 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。
// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。
// 示例 1:
// 输入:
// [[1,1,0],
//  [1,1,0],
//  [0,0,1]]
// 输出: 2
// 说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
// 第2个学生自己在一个朋友圈。所以返回2。
// 示例 2:
// 输入:
// [[1,1,0],
//  [1,1,1],
//  [0,1,1]]
// 输出: 1
// 说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
// 注意：
// N 在[1,200]的范围内。
// 对于所有学生，有M[i][i] = 1。
// 如果有M[i][j] = 1，则有M[j][i] = 1。
package main

// dfs + 数组过滤
func findCircleNum(M [][]int) int {
	visited := make([]int, len(M))
	count := 0
	// 从i开始,传i,根据i的值改j
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			dfs(M, visited, i)
			count++
		}
	}
	return count
}

func dfs(M [][]int, visited []int, i int) {
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			dfs(M, visited, j)
		}
	}
}

func findCircleNum2(M [][]int) int {
	visited := make([]int, len(M))
	queue := []int{}
	count := 0
	for i := 0; i < len(M); i++ {
		if visited[0] == 0 {
			queue = append(queue, i)
			for len(queue) != 0 {
				s := queue[0]
				queue = queue[1:]
				visited[s] = 1
				for j := 0; j < len(M); j++ {
					if M[s][j] == 1 && visited[j] == 0 {
						queue = append(queue, j)
					}
				}
			}
			count++
		}
	}
	return count
}

// 并查集
func findCircleNum3(M [][]int) int {
	parent := make([]int, len(M))

	// 重置所有根结点
	for i, _ := range parent {
		parent[i] = -1
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && i != j {
				union(parent, i, j)
			}
		}
	}
	count := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			count++
		}
	}
	return count

}

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int, x, y int) {
	xs := find(parent, x)
	ys := find(parent, y)
	if xs != ys {
		parent[xs] = ys
	}
}
