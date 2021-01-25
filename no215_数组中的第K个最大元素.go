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
	"math/rand"
	"sort"
)

func main() {
	a := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest3(a, 2))

}

// 排序
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]

}

// 类似快速排序,因为知道k在哪边,省却了一边的递归
func findKthLargest2(nums []int, k int) int {
	// 在这里要换target， 第K大是倒数
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	pivot := rand.Intn(r-l) + l
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
//为了节约空间复杂度。即 k 较小的时候使用最小堆，k 较大的时候使用最大堆
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
	return len((h))
}

func (h IntHeap) Less(i, j int) bool {
	return (h)[i] < (h)[j]
}

func (h IntHeap) Swap(i, j int) {
	(h)[i], (h)[j] = (h)[j], (h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}

//作者：xilepeng
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/dui-python3-by-xilepeng/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func findKthLargest4(nums []int, k int) int {
	nums = heapSort(nums)
	return nums[k-1]
}

func heapSort(nums []int) []int {
	lens := len(nums)
	// 建堆	lens/2后面都是叶子节点，不需要调整down()
	//  从底向上
	for i := lens / 2; i >= 0; i-- {
		down(nums, i, lens)
	}
	//将小根堆堆顶排到切片末尾(降序)
	for i := lens - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		lens--
		down(nums, 0, lens)
	}
	return nums
}
func down(nums []int, i, lens int) { //小根堆
	min := i                    //i父节点
	left, right := 2*i+1, 2*i+2 //左，右孩子
	if left < lens && nums[left] < nums[min] {
		min = left
	}
	if right < lens && nums[right] < nums[min] {
		min = right
	}
	if min != i {
		nums[min], nums[i] = nums[i], nums[min]
		down(nums, min, lens)
	}
}
