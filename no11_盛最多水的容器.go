package main

// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例:
// 输入: [1,8,6,2,5,4,8,3,7]
// 输出: 49

func main() {
	arr := []int{2, 3, 4, 5, 18, 17, 6}
	println(maxArea2(arr))

}

// 暴力计算每种围成情况的面积
func maxArea(height []int) int {
	var res int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			res = max(min(height[i], height[j])*(j-i), res)
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 双指针法
// 两个指针指向头尾,每次计算当前围城的面积之后移动
// 每次短的边向长的边的方向移动一位,再重新计算
// 注意什么时候用指针位置,什么时候用指针指向的值
func maxArea2(height []int) int {
	var res, left, right int
	right = len(height) - 1
	for left < right {
		res = max((right-left)*min(height[left], height[right]), res)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return res
}
