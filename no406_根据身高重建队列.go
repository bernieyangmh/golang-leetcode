package main

import (
	"fmt"
	"sort"
)

//假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。
//注意：
//总人数少于1100人。
//示例
//输入:
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//输出:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

func main() {
	//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

	a := [][]int{}
	a = append(a, []int{7, 0})
	a = append(a, []int{4, 4})
	a = append(a, []int{7, 1})
	a = append(a, []int{5, 0})
	a = append(a, []int{6, 1})
	a = append(a, []int{5, 2})
	fmt.Println(reconstructQueue(a))
}

// 根据身高和前面人的个数排序，然后插入到指定位置
func reconstructQueue(people [][]int) [][]int {
	sortQueue(people, 0, len(people)-1)

	res := make([][]int, len(people))

	// 一开始在这里加了太多逻辑，排序的时候条件加上[1](前面人的个数)  就能无脑插入
	for _, p := range people {
		index := p[1]
		res = append(res, []int{})
		copy(res[index+1:], res[index:])
		res[index] = p
	}

	// 最后一堆空数组不需要
	return res[:len(people)]
}

func sortQueue(people [][]int, start, end int) {
	if end-start < 1 {
		return
	}
	pivot := people[end]
	splitPoint := start

	// 类似快排， 优先考虑身高，等于的时候位置前的在前面
	for i := start; i < end; i++ {
		if people[i][0] > pivot[0] || (people[i][0] == pivot[0] && people[i][1] < pivot[1]){
			temp := people[splitPoint]
			people[splitPoint] = people[i]
			people[i] = temp
			splitPoint++
		}
	}
	people[end], people[splitPoint] = people[splitPoint], pivot
	sortQueue(people, start, splitPoint-1)
	sortQueue(people, splitPoint+1, end)
	return
}


//别人写的
func reconstructQueue2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})

	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1] +1:], result[info[1]:])
		result[info[1]] = info
	}

	return result
}