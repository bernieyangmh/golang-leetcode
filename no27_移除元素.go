package main

//1.在原数组上移除元素
//2.返回之后的长度

// 注意一下删掉元素之后要i--
// 注意i+1的边界判断
func removeElement(nums []int, val int) int {
	cnt := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			cnt--
			if i == len(nums)-1 {
				nums = nums[:i]
			} else {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}
	return cnt
}
