package main

import (
	"fmt"
)

func main()  {
	//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

	a := [][]int{}
	a = append(a, []int{7,0})
	a = append(a, []int{4,4})
	a = append(a, []int{7,1})
	a = append(a, []int{5,0})
	a = append(a, []int{6,1})
	a = append(a, []int{5,2})
	fmt.Println(a)
	sortQueue(a,0, len(a)-1)
	fmt.Println(a)
}




// 二维数组的快速排序
func sortQueue(people [][]int, start, end int){
	if end - start < 1{
		return
	}
	pivot := people[end]
	splitPoint := start


	for i:=start;i<end;i++{
		if people[i][0] > pivot[0] {
			temp := people[splitPoint]
			people[splitPoint] = people[i]
			people[i] = temp
			splitPoint++
			fmt.Println(splitPoint, people)
		}
	}
	people[end], people[splitPoint] = people[splitPoint], pivot
	sortQueue(people, start, splitPoint-1)
	sortQueue(people, splitPoint+1, end)
	return
}

