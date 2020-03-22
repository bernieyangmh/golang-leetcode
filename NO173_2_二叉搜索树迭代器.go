package main

type BSTIterator struct {
	Stack []*TreeNode
}

// 初始化一次性把所有左结点放进栈，next的时候抛出并将其右结点放进栈
func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{}
	b.Stack = []*TreeNode{}
	b.inorder(root)
	return b
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	topNode := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	if topNode.Right != nil {
		this.inorder(topNode.Right)
	}
	return topNode.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}

func (this *BSTIterator) inorder(root *TreeNode) {
	for root != nil {
		this.Stack = append(this.Stack, root)
		root = root.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
