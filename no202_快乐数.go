package main


func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)
		if slow == fast {
			break
		}
	}
	return slow == 1

}

func bitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		bit := n%10
		sum+=bit*bit
		n = n / 10
	}
	return sum

}


