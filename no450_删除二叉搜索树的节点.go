package main

//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
//一般来说，删除节点可分为两个步骤：
//首先找到需要删除的节点；
//如果找到了，删除它。
//说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 1。A 是末端节点，两个子节点都为空，直接删除
	// 2。A 只有一个非空子节点，那让这个孩子接替自己的位置
	if root.Val == key {

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 3。A 有两个子节点 必须找到左子树中最大的那个节点，或者 右子树中最小的那个节点来代替自己。 选择一种方案即可
		// 这里选择右子树最小节点
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)


	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root

}

func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// 删除左
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}


		Node := getMax(root.Left)
		root.Val = Node.Val
		root.Left = deleteNode(root.Left, Node.Val)


	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root

}

func getMax(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}
