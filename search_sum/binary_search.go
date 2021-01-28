package search_sum

// 二分查找， 普通 ｜ 左边界 ｜ 右边界

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

// 左边界
func left_bound(num []int, target int) int {
	left := 0
	right := len(num) - 1

	for left <= right {
		mid := left + (right-left)/2
		if num[mid] < target {
			left = mid + 1
		} else { // mid>=target //始终收缩右边界，才找到左边界
			right = mid - 1
		}
	}
	if left >= len(num) || num[left] != target {
		return -1
	}

	return left
}

// 右边界
func right_bound(num []int, target int) int {
	left := 0
	right := len(num)

	for left < right {
		mid := left + (right-left)/2
		if num[mid] <= target {
			left = mid + 1 // nums[mid] == target时，不要立即返回，而是增大「搜索区间」的下界left，使得区间不断向右收缩，达到锁定右侧边界的目的
		} else { // mid>target
			right = mid
		}
	}
	if left == 0 || num[left-1] == target {

		return -1
	}
	return left - 1
}


//ref : https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485044&idx=1&sn=e6b95782141c17abe206bfe2323a4226&chksm=9bd7f87caca0716aa5add0ddddce0bfe06f1f878aafb35113644ebf0cf0bfe51659da1c1b733&scene=21#wechat_redirect