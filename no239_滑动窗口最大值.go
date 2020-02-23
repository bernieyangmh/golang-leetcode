package main


//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。
//示例:
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
// 
//提示：
//你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

//暴力，超时
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := []int{nums[0]}
	windowNum := make(map[int]int)
	windowArr := []int{}


	for i := 0; i<k;i++ {
		if _, ok := windowNum[nums[i]]; ok {
			windowNum[nums[i]]+=1
		} else {
			windowNum[nums[i]] = 1
		}
		windowArr = append(windowArr, nums[i])
		res[0] = max(res[0], nums[i])
	}

	for i :=k; i<len(nums); i++{
		// 放进windows
		if _, ok := windowNum[nums[i]]; ok {
			windowNum[nums[i]]+=1
		} else {
			windowNum[nums[i]] = 1
		}
		windowArr = append(windowArr, nums[i])
		windowNum[windowArr[0]]--
		windowArr = windowArr[1:]

		curMax := windowArr[0]
		for num, cnt := range windowNum {
			if cnt > 0 {
				curMax = max(curMax, num)
			}
		}
		res = append(res, curMax)

	}
	return res

}


func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}


func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || k == 0{
		return []int{}
	}
	size := len(nums)
	maxIndex:=0
	queue := []int{}

	if k == 1 {
		return nums
	}


	for i:=0; i<k;i++{
		queue = cleanDeque(nums, queue, i, k)
		queue = append(queue, i)
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	res := make([]int, size-k+1)
	res[0] = nums[maxIndex]

	for i:=k;i<size;i++{
		queue= cleanDeque(nums, queue, i, k)
		queue = append(queue, i)
		res[i-k+1] = nums[queue[0]]
	}
	return res
}

func cleanDeque(nums, queue[]int, i, k int) []int {
	if len(queue) > 0 && queue[0] == i-k{
		queue = queue[1:]
	}
	for len(queue) > 0 && nums[i]>nums[queue[len(queue)-1]]{
		queue = queue[:len(queue)-1]
	}

	return queue
}