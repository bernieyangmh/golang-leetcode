package main

import "fmt"

// 给定一个没有重复数字的序列，返回其所有可能的全排列。
// 示例:
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func main() {
	a := []int{1, 2, 31, 4, 2, 6, 7, 18, 9, 10}

	fmt.Println(permute(a))

}

// 详解 https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/

// 和no39,40原理类似, 以nums的每个元素做根结点,叶子结点是nums的所有元素
// 已经走过的结点在下一层要从树中去掉, map[index:bool] 或arr[bool,bool...] 或位运算来过滤
// 注意回溯,返回上一个结点,再之后的路径
// 深度和nums长度一致,表明走到头,把走过的整个路径放进结果中

// ps:1.一开始想用map来排除走的路径,go语法有点难受, 位运算更好,注意要加uint32()
//    2.res要用指针,不然保留不到结果
//    3.path注意深拷贝,不然之后改动path会影响当前的结果
func permute(nums []int) [][]int {
	var res [][]int
	var size int
	size = len(nums)
	dfs(nums, size, 0, 0, []int{}, &res)
	return res
}


//todo 1。给定十个无序正整数x<=100进行排序，要求排序后每连续五个数的和不大于100

func dfs(nums []int, size, depth, used int, path []int, res *[][]int) {
	// 走到底,路径放进结果中
	if depth == size {
		// 深拷贝
		tmp := append([]int(nil), path...)
		*res = append(*res, tmp)
		return
	}
	////  //todo 1。新题目 我们只要一个结果
	//if len(*res) > 0 {
	//	return
	//}
	// 父结点下面以nums的所有元素当作叶子结点
	for i := 0; i < size; i++ {
		// used右移动i个,代表当前的叶子结点,和1作与运算 1是已经走过,0是未走过
		if (used>>uint32(i))&1 == 0 {
			// 将used的第i位变位1, 1左移动i位变成1000xxx.., 和used作异或运算
			used = used ^ (1 << uint32(i))
			// 将当前结点的值放进路径中,代表我要从这个结点继续走
			path = append(path, nums[i])
			//// todo 1.检查后5个和是否大于100
			//if len(path) >= 5 && !checkSum(path[len(path)-5:]) {
			//	return
			//}
			dfs(nums, size, depth+1, used, path, res)
			// 回溯 为走其他结点还原成还在父结点的状态; 将对应位改成0, 删掉放进的path的最后一位, 
			used = used ^ (1 << uint32(i))
			if len(path) > 0 {
				path = path[0 : len(path)-1]
			}
		}
	}
}

func checkSum(arr []int) bool {
	count := 0
	for _, v := range arr {
		count += v
		if count > 100 {
			return false
		}
	}
	return true
}
