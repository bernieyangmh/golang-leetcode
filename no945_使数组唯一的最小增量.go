package main

import (
	"sort"
)

//给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。
//返回使 A 中的每个值都是唯一的最少操作次数。
//示例 1:
//输入：[1,2,2]
//输出：1
//解释：经过一次 move 操作，数组将变为 [1, 2, 3]。
//示例 2:
//输入：[3,2,1,2,1,7]
//输出：6
//解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
//可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。

// 暴力，超时
func minIncrementForUnique(A []int) int {
	cnt := 0
	hasMap := make(map[int]int)
	for _, n := range A {
		if _, ok := hasMap[n]; !ok {
			hasMap[n] = 1
		} else {
			cur := n
			for {
				_, ok := hasMap[cur]
				if !ok {
					hasMap[cur] = 1
					break
				}
				cur++
				cnt++
			}
		}
	}
	return cnt
}

// 先排序，如果当前数比前一个小，+到前一个数+1
func minIncrementForUnique2(A []int) int {
	sort.Ints(A)
	move := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			pre := A[i]
			A[i] = A[i-1] + 1
			move += A[i] - pre
		}
	}
	return move
}

func minIncrementForUnique3(A []int) int {
	count := make([]int, 40001)
	maxNum := -1
	for _, n := range A {
		count[n]++
		maxNum = max(maxNum, n)
	}
	move := 0
	for n := 0; n <= maxNum; n++ {
		// >1 说明重复， cnt-1个该数 都+1
		if count[n] > 1 {
			d := count[n] - 1
			move += d
			count[n+1] += d
		}
	}
	// 判断最大的数有无重复，n-1个重复的直接求和得操作次数
	d := count[maxNum+1] - 1
	move += (1 + d) * d / 2
	return move
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

//路径压缩
func minIncrementForUnique4(A []int) int {
	pos := []int{}
	//最坏情况是40000个40000
	for i := 80000; i >= 0; i-- {
		pos = append(pos, -1)
	}
	move := 0
	for _, n := range A {
		b := findPos(n, pos)
		move += b - n
	}
	return move
}

func findPos(a int, pos []int) int {
	b := pos[a]
	// 没有被占用
	if b == -1 {
		pos[a] = a
		return a
	}
	//看看下一位有没有占用
	b = findPos(b+1, pos)
	// 这个位置指向下一个空位
	pos[a] = b
	return b
}
