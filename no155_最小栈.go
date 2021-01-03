// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) -- 将元素 x 推入栈中。
// pop() -- 删除栈顶的元素。
// top() -- 获取栈顶元素。
// getMin() -- 检索栈中的最小元素。
// 示例:
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.


package main

import (
	"fmt"
)

// ["MinStack","push","push","top","getMin","pop","getMin","top"]
// [[],[1],[2],[],[],[],[],[]]

func main() {
	ms := Constructor()
	println(ms.GetMin())
	ms.Push(1)
	ms.Push(2)
	println(ms.Top())
	println(ms.GetMin())
	ms.Pop()
	println(ms.GetMin())
	println(ms.Top())

}

type MinStack struct {
	Value    []int
	MinValue int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

func (this *MinStack) Push(x int) {
	// 这里一定要带上判断长度，golang int 默认0，但最小值不一定是0
	if len(this.Value) == 0 {
		this.MinValue = x
		this.Value = append(this.Value, 0)
		return
	}
	//和最小值的差
	this.Value = append(this.Value, x-this.MinValue)
	if x < this.MinValue {
		this.MinValue = x
	}
	fmt.Println(this.Value, this.MinValue)
}

func (this *MinStack) Pop() {
	if len(this.Value) == 0 {
		return
	}
	// fmt.Println(this.Value)
	if this.Value[len(this.Value)-1] < 0 {
		this.MinValue -= this.Value[len(this.Value)-1]
		this.Value = this.Value[:len(this.Value)-1]
	} else {
		this.Value = this.Value[:len(this.Value)-1]
	}

}

func (this *MinStack) Top() int {
	if len(this.Value) == 0 {
		return 0
	}
	if this.Value[len(this.Value)-1] < 0 {
		return this.MinValue
	} else {
		return this.MinValue + this.Value[len(this.Value)-1]
	}
}

func (this *MinStack) GetMin() int {
	return this.MinValue
}
