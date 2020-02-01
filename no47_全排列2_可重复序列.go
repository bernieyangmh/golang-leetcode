package main

import (
	"fmt"
	"sort"
)

// 给定一个可包含重复数字的序列，返回所有不重复的全排列。
// 输入: [1,1,2]
// 输出:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

func main() {
	a := []int{3, 3, 0, 3}

	fmt.Println(permuteUnique(a))

}

// 和46类似,唯一不同是要过滤重复元素
// 为了过滤重复元素,我们需要比较数组是否不同,在一开始就对候选数组排序比较简单
// 判断是否进入某一叶子结点前,和No.46相比多了一步判断, 该叶子结点不能和前一个叶子结点值相同,且前一个叶子结点是未走过的/走过的

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var size int
	size = len(nums)
	sort.Ints(nums)
	dfs(nums, size, 0, 0, []int{}, &res)
	return res
}

func dfs(nums []int, size, depth, used int, path []int, res *[][]int) {
	// 走到底,路径放进结果中
	if depth == size {
		// 深拷贝
		tmp := append([]int(nil), path...)

		// 这样写慢一些
		// tmp := make([]int, len(path))
		// copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	// prev := -1<<31 prev 是最小值或最大值,如果候选数组里nums有最小值或最大值不太好处理, 还是用used判断比较好
	// 父结点下面以nums的所有元素当作叶子结点
	for i := 0; i < size; i++ {

		// used右移动i个,代表当前的叶子结点,和1作与运算 1是已经走过,0是未走过
		if (used>>uint32(i))&1 == 0 {
			// 和前一个元素不等(数组已排序) 且 前一个未走过
			// 相同元素,未走过是顺序排列 1,1',1''; 走过是倒序1'',1',1
			// 如果不加(used>>uint32(i-1))&1 == 0[1], 1, 1, 1, 只能放进去1个1, 整个树无法走到底
			// 也可以用一个prev 上一次搜索的起点,遇到就跳过
			// if nums[i] == prev {continue}
			if i > 0 && nums[i] == nums[i-1] && (used>>uint32(i-1))&1 == 0 {
				continue
			}
			// 将used的第i位变位1, 1左移动i位变成1000xxx.., 和used作异或运算
			used = used ^ (1 << uint32(i))

			// 将当前结点的值放进路径中,代表我要从这个结点继续走
			path = append(path, nums[i])
			// 深度要加1
			dfs(nums, size, depth+1, used, path, res)

			// 回溯 为走其他结点还原成还在父结点的状态; 将对应位改成0, 删掉放进的path的最后一位,
			used = used ^ (1 << uint32(i))
			if len(path) > 0 {
				path = path[0 : len(path)-1]
			}
			// prev = nums[i]
		}
	}
}
