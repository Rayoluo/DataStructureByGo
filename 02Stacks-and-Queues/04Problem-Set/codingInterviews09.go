package main

import "fmt"

// 双栈实现队列
type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	// 队列为空的情况
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return -1
	}
	if len(this.stack1) != 0 && len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			value := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, value)
		}
	}
	ret := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return ret
}

func main() {
	queue := Constructor()
	queue.AppendTail(3)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
	queue.AppendTail(4)
	queue.AppendTail(5)
	fmt.Println(queue.DeleteHead())
	queue.AppendTail(6)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())

}
