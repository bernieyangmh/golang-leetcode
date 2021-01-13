package collections


func maxWater( arr []int ) int64 {
	// write code here

	n := len(arr)
	if n < 2{
		return 0
	}

	left, right  := 0, n-1
	l_max, r_max := arr[0], arr[n-1]
	ans := 0

	for left <= right {
		l_max = max(l_max, arr[left])
		r_max = max(r_max, arr[right])

		if l_max < r_max {
			ans += l_max-arr[left]
			left++
		} else {
			ans += r_max-arr[right]
			right --
		}
	}
	return int64(ans)

}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
