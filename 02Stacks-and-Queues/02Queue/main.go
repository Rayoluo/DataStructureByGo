package main

import (
	"fmt"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/02Queue/arrayQueue"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/02Queue/queue"
)

func main() {
	var queue queue.Queue
	queue = arrayQueue.NewArrayQueue(20)
	fmt.Println(queue.ArrayQueueToString())
	for i := 0; i < 10; i++ {
		if i%3 == 1 {
			queue.Dequeue()
		}
		queue.Enqueue(i)
		fmt.Println(queue.ArrayQueueToString())
	}
}
