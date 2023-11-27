package min_stack

import "math"

type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	s := MinStack{}
	s.min = append(s.min, math.MaxInt)
	return s
}

func (this *MinStack) Push(val int) {
	m := this.min[len(this.min)-1]
	if val < m {
		m = val
	}
	this.data = append(this.data, val)
	this.min = append(this.min, m)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
