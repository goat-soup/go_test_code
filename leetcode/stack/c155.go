package stack

import "fmt"

type MinStack struct {
	v     []int
	size  int
	min_v int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.size++
	this.v = append(this.v, val)
	this.min_v = min(this.v)
}

func (this *MinStack) Pop() {
	if this.size == 0 {
		return
	}
	this.size--
	this.v = this.v[0:this.size]
	this.min_v = min(this.v)
}

func (this *MinStack) Top() int {
	return this.v[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min_v
}

func min(v []int) int {
	if len(v) == 0 {
		return 0
	}
	a := v[0]
	for i := 1; i < len(v); i++ {
		if a > v[i] {
			a = v[i]
		}
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func Mystack() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
