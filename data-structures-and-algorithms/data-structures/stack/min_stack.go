package stack

type MinStack struct {
	items []int
	minItems []int
}


func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	this.items = append(this.items, val)
	if len(this.minItems) == 0 {
		this.minItems = append(this.minItems, val)
	} else {
		lastIdx := len(this.minItems) - 1
		minMin := this.minItems[lastIdx]
		this.minItems = append(this.minItems, min(val,minMin))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop()  {
	lastIdx := len(this.items) - 1
	this.items = this.items[:lastIdx]
	this.minItems = this.minItems[:lastIdx]
}


func (this *MinStack) Top() int {
	lastIdx := len(this.items) - 1
	return this.items[lastIdx]
}


func (this *MinStack) GetMin() int {
	lastIdx := len(this.items) - 1
	return this.minItems[lastIdx]
}
