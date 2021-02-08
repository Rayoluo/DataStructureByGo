package main

import "fmt"

func main() {
	maxq := ConstructorForMaxQueue()
	fmt.Println(maxq.Pop_front())
	fmt.Println(maxq.Pop_front())
	fmt.Println(maxq.Pop_front())
	fmt.Println(maxq.Pop_front())
	fmt.Println(maxq.Pop_front())
	maxq.Push_back(15)
	fmt.Println(maxq.Max_value())
	maxq.Push_back(9)
	fmt.Println(maxq.Max_value())
}

type MaxQueue struct {
	dataQueue []int // 普通队列
	deque     []int // 双端队列
}

func ConstructorForMaxQueue() MaxQueue {
	return MaxQueue{
		dataQueue: make([]int, 0),
		deque:     make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.dataQueue = append(this.dataQueue, value)
	i := len(this.deque) - 1
	for i >= 0 {
		if this.deque[i] > value {
			break
		}
		i--
	}
	this.deque = append(this.deque[:i+1], value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.dataQueue) == 0 {
		return -1
	}
	ret := this.dataQueue[0]
	if this.dataQueue[0] == this.deque[0] {
		this.deque = this.deque[1:]
	}
	this.dataQueue = this.dataQueue[1:]
	return ret
}
