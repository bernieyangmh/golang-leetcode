package collections

func countbit(num int) int {
	count := 0
	for num != 0 {
		//不停右移， 和1与操作
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
