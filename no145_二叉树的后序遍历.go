package main

//给定一个二叉树，返回它的 后序 遍历。
//示例:
//输入: [1,null,2,3]
//1
// \
//  2
// /
//3
//输出: [3,2,1]

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	// 防止再进入走过的子结点
	passMap := make(map[*TreeNode]bool)
	nodeStack := []*TreeNode{root}
	for len(nodeStack) > 0 {
		// 有左结点push左结点
		top := nodeStack[len(nodeStack)-1]
		// fmt.Println(top)
		if top.Left != nil {
			if _, ok := passMap[top.Left]; !ok {
				nodeStack = append(nodeStack, top.Left)
				continue
			}
		}
		if top.Right != nil {
			if _, ok := passMap[top.Right]; !ok {
				nodeStack = append(nodeStack, top.Right)
				continue
			}
		}
		// 左右结点为nil或均已走过 抛出顶点
		res = append(res, top.Val)
		passMap[top] = true
		nodeStack = nodeStack[:len(nodeStack)-1]
	}
	return res
}

//puh两次，先取一次判断是否和之后的是一样的
//一样说明是第一次进该结点，把右左子结点pish进去
//不一样说明是出该结点，把结点值放进结果里
func postorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == nil {
			continue
		}
		if len(stack) > 0 && cur == stack[len(stack)-1] {
			stack = append(stack, cur.Right)
			stack = append(stack, cur.Right)
			stack = append(stack, cur.Left)
			stack = append(stack, cur.Left)
		} else {
			res = append(res, cur.Val)
		}
	}
	return res
}
