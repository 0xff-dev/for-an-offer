package MinStack


type MinStack struct {
	nums []int
	min  []int
	size int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), -1}
}


func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (this *MinStack) Push(x int)  {
	this.nums = append(this.nums, x)
	if this.size == -1 {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, Min(x, this.min[this.size]))
	}
	this.size ++
}


func (this *MinStack) Pop()  {
	if this.size > -1 {
		this.nums = this.nums[:this.size]
		this.min = this.min[:this.size]
		this.size --
	}
}


func (this *MinStack) Top() int {
	if this.size > -1 {
		return this.nums[this.size]
	}
	return -1
}


func (this *MinStack) GetMin() int {
	if this.size > -1 {
		return this.min[this.size]
	}
	return -1
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
