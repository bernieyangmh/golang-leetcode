package search_sum

func binarySearch(num []int, target int) int {
	left := 0
	right := len(num) - 1

	for left <= right {
		mid := left + (right-1)/2
		if num[mid] == target {
			return mid
		}
		if num[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearch2(num []int, target int) int {
	left := 0
	right := len(num) - 1

	for left < right {
		mid := (left + right + 1) >> 1
		if target < num[mid] {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if num[left] == target {
		return left
	}
	return -1
}
