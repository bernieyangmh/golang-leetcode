package main

//请实现有重复数字的升序数组的二分查找。
//输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。

func upper_bound_(n int, v int, a []int) int {
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if a[mid] < v {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r + 1
}

func binaryFind(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
