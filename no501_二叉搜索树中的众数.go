// 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
// 假定 BST 有如下定义：
// 结点左子树中所含结点的值小于等于当前结点的值
// 结点右子树中所含结点的值大于等于当前结点的值
// 左子树和右子树都是二叉搜索树
// 例如：
// 给定 BST [1,null,2,2],

//    1
//     \
//      2
//     /
//    2
// 返回[2].
// 提示：如果众数超过1个，不需考虑输出顺序
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//暴力遍历,放进数组里,最后用map统计数量,然后求众数
func findMode(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return []int{}
	}

	arr = inorder(root, arr)

	aidNum := 0
	modeMap := make(map[int]int)
	res := []int{}

	for i := 0; i < len(arr); i++ {
		if _, ok := modeMap[arr[i]]; ok {
			modeMap[arr[i]]++
		} else {
			modeMap[arr[i]] = 1
		}
	}
	for num, cnt := range modeMap {
		if cnt > aidNum {
			res = []int{num}
			aidNum = cnt
		} else if cnt == aidNum {
			res = append(res, num)
		}
	}

	return res
}

func inorder(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = inorder(root.Left, arr)
	arr = append(arr, root.Val)
	arr = inorder(root.Right, arr)

	return arr
}

// 递归时判断,参数多但不需要重复多次
func findMode2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res, _, _, _ = inorder2(root, res, 1<<32, -1, -1)

	return res
}

// 结点, 结果, 上一个数的数值, 上一个数的数量, 当前符合众数的值
// 一定要记住传进去的参数也要返回,不然回退到父节点并不是子结点已经计算好的值
func inorder2(root *TreeNode, arr []int, num, cnt, aidcnt int) ([]int, int, int, int) {
	if root == nil {
		return arr, num, cnt, aidcnt
	}
	arr, num, cnt, aidcnt = inorder2(root.Left, arr, num, cnt, aidcnt)

	// 如果结点值和上一个值不同, 更新为结点值,数量为1
	// 如果相同,数量+1
	if root.Val != num {
		num = root.Val
		cnt = 1
	} else {
		cnt++
	}

	// 如果数量大于当前众数的标准,重置结果
	// 如果数量等于当前众数的标准,新加一个
	if cnt > aidcnt {
		arr = []int{num}
		aidcnt = cnt

	} else if cnt == aidcnt {
		arr = append(arr, num)
	}

	arr, num, cnt, aidcnt = inorder2(root.Right, arr, num, cnt, aidcnt)

	return arr, num, cnt, aidcnt
}
