package main

import "fmt"

type MinStack struct {
	stack1 []int // 数据栈
	stack2 []int // 辅助栈
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack2) == 0 {
		this.stack2 = append(this.stack2, x)
	} else {
		minVal := this.stack2[len(this.stack2)-1]
		if x < minVal {
			this.stack2 = append(this.stack2, x)
		} else {
			this.stack2 = append(this.stack2, minVal)
		}
	}
	this.stack1 = append(this.stack1, x)
}

func (this *MinStack) Pop() {
	if len(this.stack1) == 0 {
		panic("Error! Empty Stack!")
	}
	this.stack1 = this.stack1[:len(this.stack1)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack1) == 0 {
		panic("Error! Empty Stack!")
	}
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack2) == 0 {
		panic("Error! Empty Stack!")
	}
	return this.stack2[len(this.stack2)-1]
}

func main() {
	minStack := Constructor1()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.Min())
}
