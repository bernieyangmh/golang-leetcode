package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	CurIndex int
	ValArr   []int
	Size     int
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{CurIndex: 0}
	b.dfs(root)
	b.Size = len(b.ValArr)
	return b
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.CurIndex++
	return this.ValArr[this.CurIndex-1]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.CurIndex < this.Size
}

func (this *BSTIterator) dfs(root *TreeNode) {
	if root == nil {
		return
	}
	this.dfs(root.Left)
	this.ValArr = append(this.ValArr, root.Val)
	this.dfs(root.Right)
	return
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
