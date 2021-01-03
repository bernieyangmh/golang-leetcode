package main

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
	Value  []int
	Helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {

	this.Value = append(this.Value, x)
	// 辅助栈 放入当前值或原栈顶值。 即此刻的最小值
	if len(this.Helper) == 0 || x <= this.Helper[len(this.Helper)-1] {
		this.Helper = append(this.Helper, x)
	} else {
		this.Helper = append(this.Helper, this.Helper[len(this.Helper)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.Value) > 0 {
		this.Value = this.Value[:len(this.Value)-1]
		this.Helper = this.Helper[:len(this.Helper)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Value) > 0 {
		return this.Value[len(this.Value)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.Helper) > 0 {
		return this.Helper[len(this.Helper)-1]
	}
	return 0
}
