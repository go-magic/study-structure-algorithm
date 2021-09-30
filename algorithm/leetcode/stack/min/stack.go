package stack

type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minData) == 0 {
		this.minData = append(this.minData, x)
		return
	}
	minData := this.minData[len(this.minData)-1]
	if x < minData {
		this.minData = append(this.minData, x)
		return
	}
	//相同的值入栈是为了防止两个相同的数导致临时栈被弹出
	this.minData = append(this.minData, minData)
}

func (this *MinStack) Pop() {
	length := len(this.data)
	if length == 0 {
		return
	}
	this.data = this.data[:length-1]
	this.minData = this.minData[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	length := len(this.data)
	if length == 0 {
		return 0
	}
	return this.data[length-1]
}

func (this *MinStack) Min() int {
	length := len(this.minData)
	if length == 0 {
		return 0
	}
	return this.minData[length-1]
}
