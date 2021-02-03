package arrayQueue

import (
	"bytes"
	"fmt"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/02Queue/array"
)

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array.NewArray(capacity),
	}
}

type ArrayQueue struct {
	*array.MyArray
}

func (aq *ArrayQueue) GetQueueSize() int {
	return aq.GetSize()
}

func (aq *ArrayQueue) IsQueueEmpty() bool {
	return aq.IsEmpty()
}

func (aq *ArrayQueue) Enqueue(e interface{}) {
	aq.AddLast(e)
}

func (aq *ArrayQueue) Dequeue() interface{} {
	return aq.DeleteFirst()
}

func (aq *ArrayQueue) GetFront() interface{} {
	return aq.Get(0)
}

func (aq *ArrayQueue) ArrayQueueToString() string {
	var buf bytes.Buffer
	buf.WriteString("Queue:")
	if aq.GetSize() == 0 {
		buf.WriteString("[] tail")
	} else {
		buf.WriteString("[")
		for i := 0; i < aq.GetSize(); i++ {
			buf.WriteString(fmt.Sprintf("%v", aq.Get(i)))
			if i != aq.GetSize()-1 {
				buf.WriteString(" ,")
			}
		}
		buf.WriteString("] tail")
	}
	return buf.String()
}
