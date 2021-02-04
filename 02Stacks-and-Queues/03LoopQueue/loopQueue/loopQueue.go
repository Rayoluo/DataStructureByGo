package loopQueue

import (
	"bytes"
	"fmt"
)

type LoopQueue struct {
	arr         []interface{}
	front, tail int
	size        int
}

func NewLoopQueue(capacity int) *LoopQueue {
	return &LoopQueue{
		arr: make([]interface{}, capacity+1),
	}
}

func (lq *LoopQueue) GetQueueSize() int {
	return lq.size
}

func (lq *LoopQueue) GetCapacity() int {
	return len(lq.arr) - 1
}

// 如果front=tail则队列为空
func (lq *LoopQueue) IsQueueEmpty() bool {
	return lq.front == lq.tail
}

// 入队
func (lq *LoopQueue) Enqueue(e interface{}) {
	if (lq.tail+1)%len(lq.arr) == lq.front {
		lq.resize(lq.GetCapacity() * 2)
	}
	lq.arr[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.arr)
	lq.size++
}

func (lq *LoopQueue) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity+1)
	for i := 0; i < lq.size; i++ {
		newArr[i] = lq.arr[(lq.front+i)%len(lq.arr)]
	}
	lq.arr = newArr
	lq.front, lq.tail = 0, lq.size
}

func (lq *LoopQueue) Dequeue() (ret interface{}) {
	if lq.IsQueueEmpty() {
		panic("Dequeue Fail! Empty Queue!")
	}
	ret = lq.arr[lq.front]
	lq.arr[lq.front] = nil
	lq.front = (lq.front + 1) % len(lq.arr)
	lq.size--
	if lq.size <= lq.GetCapacity()/4 && lq.GetCapacity()/2 != 0 {
		lq.resize(lq.GetCapacity() / 2)
	}
	return
}

func (lq *LoopQueue) GetFront() interface{} {
	if lq.IsQueueEmpty() {
		panic("GetFront Fail! Queue Empty!")
	}
	return lq.arr[lq.front]
}

func (lq *LoopQueue) LoopQueueToString() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Size:%d Capacity:%d\n", lq.GetQueueSize(), lq.GetCapacity()))
	buf.WriteString("Queue:")
	if lq.GetQueueSize() == 0 {
		buf.WriteString("front [] tail")
	} else {
		buf.WriteString("front [")
		for i := lq.front; i != lq.tail; i = (i + 1) % len(lq.arr) {
			buf.WriteString(fmt.Sprintf("%v", lq.arr[i]))
			if (i+1)%len(lq.arr) != lq.tail {
				buf.WriteString(" ,")
			}
		}
		buf.WriteString("] tail")
	}
	return buf.String()
}
