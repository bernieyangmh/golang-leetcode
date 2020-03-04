package main

import (
	"math"
)

// 两个指针，指向两个数组，left和right两个数到最后表示中间的两个数
// right从两个数组排序走，left保存上一次right的数
// 总长度偶数的话，中位数就是（left+right）/2；奇就是right
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	midLen := m + n
	left, right := -1, -1
	aStart, bStart := 0, 0
	for i := 0; i <= midLen/2; i++ {
		left = right
		// 当bStart >= n; aStart++
		if aStart < m && (bStart >= n || nums1[aStart] < nums2[bStart]) {
			right = nums1[aStart]
			aStart++
		} else {
			right = nums2[bStart]
			bStart++
		}
	}
	if (midLen & 1) == 0 {
		return float64(left+right) / 2.0
	} else {
		return float64(right)
	}
}

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/
//https: //leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/he-bing-yi-hou-zhao-gui-bing-guo-cheng-zhong-zhao-/
//关键在各种边界的处理
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	// 偶数，left和right是中间的两个数；奇数，left和right相同
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2

	return float64(getKth(nums1, nums2, 0, n-1, 0, m-1, left)+getKth(nums1, nums2, 0, n-1, 0, m-1, right)) * 0.5
}

func getKth(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	//两个数组的长度
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	//始终保持len1，nums1最短
	if len1 > len2 {
		return getKth(nums2, nums1, start2, end2, start1, end1, k)
	}
	// 短的数组空了，直接返回从start起 第k小的数
	if len1 == 0 {
		return nums2[start2+k-1]
	}

	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2, start1, end1, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, nums2, i+1, end1, start2, end2, k-(i-start1+1))
	}

}

// 二分法
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	m := len(nums1)
	n := len(nums2)

	left := 0
	right := m

	for left <= right {

		//i,j代表两个数组右部分的第一个数
		//size左=(m+n+1)/2 j=(m+n+1)/2-i
		i := (left + right) >> 1
		j := ((m + n + 1) >> 1) - i
		var (
			nums1LeftMax, nums1RightMin, nums2LeftMax, nums2RightMin int
		)
		// i移动到最左边，我们取两个左部分最大值的时候一定不能取它，故取负无穷
		if i == 0 {
			nums1LeftMax = -1 << 31
		} else {
			nums1LeftMax = nums1[i-1]
		}

		//i移动到最右边，我们取两个部分右部分的最大值的时候不能取它，故取正无穷
		if i == m {
			nums1RightMin = 1<<31 - 1
		} else {
			nums1RightMin = nums1[i]
		}

		if j == 0 {
			nums2LeftMax = math.MinInt32
		} else {
			nums2LeftMax = nums2[j-1]
		}
		if j == n {
			nums2RightMin = math.MaxInt32
		} else {
			nums2RightMin = nums2[j]
		}

		//如果数组1的左边小于等于数组二的右边，且数组二的左边小于等于数组二的右边，已找到该分界线
		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			if ((m + n) & 1) == 1 {
				// 左边两个部分的最大值
				return float64(max(nums1LeftMax, nums2LeftMax))
			} else {
				// (左边两个部分的最大值，右边两个部分的最小值) / 2
				return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
			}
			// 数组2的左边大于数组1的右边
			//2 4 5 15            2 4 5 15
			//  |                     |
			//1 7 8 10 17         1 7 8 10 17
			//    |                 |
		} else if nums2LeftMax > nums1RightMin {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return 0.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
