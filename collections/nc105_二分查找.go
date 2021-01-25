package main

//请实现有重复数字的升序数组的二分查找。
//输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。

func upper_bound_( n int ,  v int ,  a []int ) int {
	l, r := 0, n
	for l < r {
		mid := l +(r-l)/2
		if a[mid] < v {
			l = mid+1
		}else {
			r = mid
		}
	}
	return r+1
}
