// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// Note:

// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
// Example:

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// Output: [1,2,2,3,5,6]

package main

import (
	"fmt"
)

func main() {
	m := []int{9, 312, 6, 12, 312, 423, 0, 0, 0, 0, 0}
	n := []int{1, 5, 3, 76, 8}

	mergech(m, 6, n, 5)
	fmt.Println(m)
}

// 从后往前插入
func merge(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {
		if nums1[m-1] <= nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n -= 1
		} else {
			nums1[m+n-1] = nums1[m-1]
			m -= 1
		}

	}
	fmt.Println(nums1, nums2, m, n)
	if n > 0 {
		for i := n - 1; i >= 0; i-- {
			nums1[i] = nums2[i]
		}

	}
}

func mergech(nums1 []int, m int, nums2 []int, n int) {
	c := make(chan int, m+n)

	for m > 0 && n > 0 {
		if nums1[m-1] <= nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			c <- nums2[n-1]
			n -= 1

		} else {
			nums1[m+n-1] = nums1[m-1]
			c <- nums1[m-1]
			m -= 1

		}

	}
	fmt.Println(nums1, nums2, m, n)
	if n > 0 {
		for i := n - 1; i >= 0; i-- {
			nums1[i] = nums2[i]
			c <- nums2[i]
		}
	}
	l := len(c)
	for i := l; i > 0; i-- {
		fmt.Println(<-c)
	}
}

// class Solution:
//     def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
//         """
//         Do not return anything, modify nums1 in-place instead.
//         """
//         # 整体思路相似，只不过没有使用 current 指针记录当前填补位置
//         while m > 0 and n > 0:
//             if nums1[m-1] <= nums2[n-1]:
//                 nums1[m+n-1] = nums2[n-1]
//                 n -= 1
//             else:
//                 nums1[m+n-1] = nums1[m-1]
//                 m -=1
//         """
//         由于没有使用 current，第一步比较结束后有两种情况:
//             1. 指针 m>0，n=0，此时不需要做任何处理
//             2. 指针 n>0，m=0，此时需要将 nums2 指针左侧元素全部拷贝到 nums1 的前 n 位
//         """
//         if n > 0:
//             nums1[:n] = nums2[:n]
