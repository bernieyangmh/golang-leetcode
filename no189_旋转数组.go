package main

//三步反转
func rotate(nums []int, k int)  {
	if len(nums) <2 {
		return
	}
	if k > len(nums) {
		k -=  len(nums)
	}
	//0   ~ n-k-1
	//n-k ~ n-1
	//0   ~ n-1
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}


func reverse(arr []int, start, end int) {
	for start < end {
		tmp := arr[start]
		arr[start] = arr[end]
		arr[end] = tmp
		start += 1
		end -= 1
	}
}


func rotate2(nums []int, k int) {
	length := len(nums)
	k = k%length

	count := 0
	for start := 0;count < length;start++{
		cur := start
		prev := nums[start]
		for {
			next := (cur+k)%length
			tmp := nums[next]
			nums[next] = prev
			prev = tmp
			cur = next
			count++
			if start == cur {
				break
			}
		}
	}
}