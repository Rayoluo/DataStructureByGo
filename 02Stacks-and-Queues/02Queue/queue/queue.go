package queue

type Queue interface {
	GetQueueSize() int
	IsQueueEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
	ArrayQueueToString() string
}
