package main

import (
	"fmt"
	"sort"
)

//给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
//然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
//你需要计算完成所有任务所需要的最短时间。
//示例 1：
//输入: tasks = ["A","A","A","B","B","B"], n = 2
//输出: 8
//执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
//注：
//任务的总个数为 [1, 10000]。
//n 的取值范围为 [0, 100]。

func main()  {
	a := "AA"
	for _,i:=range a {
		fmt.Println(i)
	}
}


func leastInterval(tasks []byte, n int) int {
	mapL := make([]int, 26)
	for _, l := range tasks {
		mapL[l-65]++
	}
	sort.Ints(mapL)
	time := 0
	for mapL[25] > 0 {
		i:=0
		for i<n{
			if mapL[25]==0 {
				break
			}
			if i < 26 && mapL[25-i] > 0 {
				mapL[25-i]--
			}
			time++
			i++
		}
		sort.Ints(mapL)
	}
	return time
}

func leastInterval2(tasks []byte, n int) int {
	task_map := make(map[byte]int)
	max_count := 0
	difference := 0

	for _, c := range tasks{
		var count int
		if _, ok := task_map[c];ok{
			count = task_map[c] + 1
		} else {
			count = 1
		}
		task_map[c] = count
		max_count = max(max_count, count)
	}

	for _, cnt := range task_map{
		if cnt == max_count{
			difference++
		}
	}
	// 最多数量A 三个组成两个间隔*间隔的长度n+1 + 最后一次排序的数量
	// A-X-X-A-X-X-A
	num1 := (n+1) * (max_count-1)+difference
	num2 := len(tasks)
	return max(num1, num2)
}


func leastInterval3(tasks []byte, n int) int {
	if len(tasks) <= 1 || n < 1 {
		return len(tasks)
	}
	mapL := make([]int, 26)
	for _, l := range tasks {
		mapL[l-65]++
	}
	sort.Ints(mapL)

	max_count := mapL[25]
	// 最多数量A 三个组成两个间隔*间隔的长度n+1 + 最后一次排序的数量
	// A-X-X-A-X-X-A
	resCount := (max_count-1)*(n+1) + 1

	//如果有任务数量等于最大数量，在后面+1
	//A-B-X-A-B-X-A-B
	i := 24
	for i>=0&&mapL[i]==max_count{
		resCount++
		i--
	}
	//当任务种类少，但是任务数量多，我们所有任务都可以执行不需等待，这是任务的数量就是计算的时间
	return max(resCount, len(tasks))
}






func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}