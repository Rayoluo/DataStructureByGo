package main

import (
	"fmt"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/01Stack/arrayStack"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/01Stack/stack"
)

func main() {
	var stack stack.Stack
	stack = arrayStack.NewArrayStack(10)
	fmt.Println(stack.ArrayStackToString())
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.ArrayStackToString())
	stack.Pop()
	fmt.Println(stack.ArrayStackToString())
}
