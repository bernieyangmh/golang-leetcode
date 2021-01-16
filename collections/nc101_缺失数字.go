package collections

func solve(a []int) int {
	// write code here
	if len(a) < 2 {
		return 0
	}

	for i := 1; i < len(a); i++ {
		// 跳过一个数判断就行了
		if a[i]-2 == a[i-1] {
			return a[i] - 1
		}
	}
	return 0
}
