package main

import (
	"fmt"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/03LoopQueue/loopQueue"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/03LoopQueue/queue"
)

func main() {
	var queue queue.Queue
	queue = loopQueue.NewLoopQueue(10)
	// fmt.Println(queue.LoopQueueToString())
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue.LoopQueueToString())
		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue.LoopQueueToString())
		}
	}
}
