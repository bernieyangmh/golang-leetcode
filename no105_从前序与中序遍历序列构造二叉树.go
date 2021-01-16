// 根据一棵树的前序遍历与中序遍历构造二叉树。
// 注意:
// 你可以假设树中没有重复的元素。
// 例如，给出
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：
//     3
//    / \
//   9  20
//     /  \
//    15   7
package main

// 前序遍历 --> 根左右
// 中序遍历 --> 左根右
// 后序遍历 --> 左右根

// 根是前序的第一个，后序的最后一个。 找到其在中序的索引i
// 递归调用
//前+中	左边 前[1:i+1] 中[:i]
//	    右边 前[i+1:] 中[i+1:]
//中+后	左边 中[:i] 后[:i]
//	    右边 中[i+1:] 后[i:len-1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序确定根,在中序中找到根值,左右分别是在根的左右结点
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	// return helper(preorder, inorder, 0, len(preorder), 0, len(inorder))
	return helper2(preorder, inorder, 0, len(preorder), 0, len(inorder), inorderMap)

}

func helper(preorder, inorder []int, ps, pe, is, ie int) *TreeNode {
	if ps == pe {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[ps]
	rootIndex := 0
	for i := is; i < ie; i++ {
		if root.Val == inorder[i] {
			rootIndex = i
			break
		}
	}
	// 前序遍历 preorder = [3,9,20,15,7]
	// 中序遍历 inorder = [9,3,15,20,7]
	//从中序剩下的开头到根所在的位置都是左边结点

	LeftNum := rootIndex - is
	// 前序跳过左边的,就是下一个根的位置     3    1		 2			0		1
	// 							        9    2		2		   0      0
	//									15   3	    3         2      2
	root.Left = helper(preorder, inorder, ps+1, ps+LeftNum+1, is, rootIndex)
	//									20   2	          5		2		5
	//									7    4            5      4         5
	root.Right = helper(preorder, inorder, ps+LeftNum+1, pe, rootIndex+1, ie)

	return root

}

// 用map存inorder
func helper2(preorder, inorder []int, ps, pe, is, ie int, inorderMap map[int]int) *TreeNode {
	if ps == pe {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[ps]
	rootIndex := inorderMap[root.Val]
	// 前序遍历 preorder = [3,9,20,15,7]
	// 中序遍历 inorder = [9,3,15,20,7]
	//从中序剩下的开头到根所在的位置都是左边结点

	LeftNum := rootIndex - is
	// 前序跳过左边的,就是下一个根的位置     3    1		 2			0		1
	// 							        9    2		2		   0      0
	//									15   3	    3         2      2
	root.Left = helper(preorder, inorder, ps+1, ps+LeftNum+1, is, rootIndex)
	//									20   2	          5		2		5
	//									7    4            5      4         5
	root.Right = helper(preorder, inorder, ps+LeftNum+1, pe, rootIndex+1, ie)

	return root
}

// 简单点
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	root_node := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	//找到根(第一个) 在中序列的位置， 也是左边的数量
	i := 0
	for ; i < len(preorder); i++ {
		if root_node.Val == inorder[i] {
			break
		}
	}
	// 左边是 先的 1～i+1， 中的前i
	root_node.Left = buildTree(preorder[1:i+1], inorder[:i])
	//先的 后i+1， 中的后i+1
	root_node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root_node
}
