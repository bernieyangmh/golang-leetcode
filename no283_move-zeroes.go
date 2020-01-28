// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 示例:
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 说明:
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

package main

import "fmt"

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)

}

// 遍历，对每一个非0数字改动到慢指针上
// 慢指针后面的全是0
func moveZeroes(nums []int) {
	var slow int
	for _, value := range nums {
		if value != 0 {
			nums[slow] = value
			slow += 1
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 慢指针永远指向该为0的元素
func moveZeroes2(nums []int) {
	var slow int
	for index, value := range nums {
		if value != 0 {
			if index != slow {
				nums[slow] = nums[index]
				nums[index] = 0
			}
			slow += 1
		}
	}
}
