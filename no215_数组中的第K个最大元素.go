// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 示例 1:
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest2(a, 2))

}

// 排序
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]

}

// 类似快速排序,因为知道k在哪边,省却了一边的递归
func findKthLargest2(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	pivot := l
	pivot = partition(nums, l, r, pivot)

	if k == pivot {
		return nums[k]
	} else if k < pivot {
		return quickSelect(nums, l, pivot-1, k)
	}
	return quickSelect(nums, pivot+1, r, k)

}

func partition(nums []int, l, r, pivot int) int {
	pivotNum := nums[pivot]
	swap(nums, pivot, r)
	s := l

	for i := l; i <= r; i++ {
		if nums[i] < pivotNum {
			swap(nums, s, i)
			s++
		}
	}
	swap(nums, s, r)
	return s

}

func swap(nums []int, a, b int) {
	tmp := nums[b]
	nums[b] = nums[a]
	nums[a] = tmp
}

// 维持一个最小堆, 大小为k,走到底之后堆顶是所求数
func findKthLargest3(nums []int, k int) int {
	q := make(IntHeap, 0, k)
	for i, n := range nums {
		if i < k {
			heap.Push(&q, n)
		} else if n > q[0] {
			heap.Pop(&q)
			heap.Push(&q, n)
		}
	}
	return q[0]
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}
