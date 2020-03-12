package main

//根据一棵树的中序遍历与后序遍历构造二叉树。
//注意:
//你可以假设树中没有重复的元素。
//例如，给出
//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//返回如下的二叉树：
//3
/// \
//9  20
///  \
//15   7

//中序 左 根 右
//后序 左 右 根
func buildTree(inorder []int, postorder []int) *TreeNode {

	if len(inorder)-1 != len(postorder)-1 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	return helper(postorder, 0, len(inorder)-1, 0, len(postorder)-1, inorderMap)
}


func helper(postorder []int, inLeft, inRight, postLeft, postRight int, inorderMap map[int]int) *TreeNode {
	if inLeft > inRight || postLeft > postRight {
		return nil
	}

	pivot := postorder[postRight]
	pivotIndex := inorderMap[pivot]
	root := new(TreeNode)
	root.Val = pivot
	//4 2 8 5 9 1 6 10 3 7
	//4 8 9 5 2 10 6 7 3 1
	// 根左左子树的区间的 前序是 [中序左边界:根在中序的索引-1]，后序 [postLeft保持不变:右边收缩跨过右子树]
	//4 2 8 5 9
	//4 8 9 5 2
	//根的右子树区间的 前序是 [根在中序的索引+1:inRight保持不变] 后序 [左边扩张跨过左子树:右边-1去掉根]
	//			  6 10 3 7
	//			10 6 7 3
	//根是后序的最右边，postLeft-inLeft+pivotIndex
	root.Left = helper(postorder, inLeft, pivotIndex-1, postLeft, postLeft-inLeft+pivotIndex-1, inorderMap)
	root.Right = helper(postorder, pivotIndex+1, inRight, postLeft-inLeft+pivotIndex, postRight-1, inorderMap)
	return root
}
