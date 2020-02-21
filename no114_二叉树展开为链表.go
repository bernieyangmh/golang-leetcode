// 给定一个二叉树，原地将它展开为链表。
// 例如，给定二叉树

//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// 将其展开为：
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// 暴力,中序遍历放进数组里,最后重新组成链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	nodeArr := []*TreeNode{}
	nodeArr = preorder(root, nodeArr)
	// 左边都不要了
	for i := 1; i < len(nodeArr); i++ {
		nodeArr[i-1].Right = nodeArr[i]
		nodeArr[i-1].Left = nil
	}
	nodeArr[len(nodeArr)-1].Left = nil
	nodeArr[len(nodeArr)-1].Right = nil
	return
}

func preorder(root *TreeNode, nodeArr []*TreeNode) []*TreeNode {
	if root == nil {
		return nodeArr
	}
	nodeArr = append(nodeArr, root)
	nodeArr = preorder(root.Left, nodeArr)
	nodeArr = preorder(root.Right, nodeArr)
	return nodeArr
}



func flatten2(root *TreeNode) {
	for root != nil {
		// 如果左边是空,跳过下一个
		if root.Left ==nil{
			root = root.Right
		} else {
			// pre代表左子树的最右端
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			// 将右子树放到左子树的最右端结点的右边
			pre.Right = root.Right
			// 整个左子树移到右子树上
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		}
	}
}



func flatten3(root *TreeNode) {
	if root == nil {
		return
	}
	s := []*TreeNode{}
	s = append(s, root)
	var pre *TreeNode
	for len(s)>0 {
		temp := s[len(s)-1]
		s = s[:len(s)-1]
		if pre != nil {
			pre.Right = temp
			pre.Left = nil
		}
		if temp.Right != nil {
			s = append(s, temp.Right)
		}
		if temp.Left != nil {
			s = append(s, temp.Left)
		}
		pre =temp
	}
}


// 右子树->左子树->根节点。
// 以右左根的顺序遍历,
func flatten4(root *TreeNode) {
	stack := []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(stack)>0{
		for root !=nil{
			// 先放根
			stack = append(stack, root)
			// 再放右结点
			root = root.Right
		}
		// 现在在最右结点
		root = stack[len(stack)-1]
		//如果没有左结点,或者左结点访问过,访问操作根结点
		if root.Left == nil || root.Left == pre {
			stack = stack[:len(stack)-1]

			// 根的右边指向它的根 6-5-4-3-2-1
			root.Right = pre
			root.Left = nil

			pre = root
			root = nil
		}else{
			// 否则左结点还没访问过就访问左结点
			root = root.Left
		}
	}
}