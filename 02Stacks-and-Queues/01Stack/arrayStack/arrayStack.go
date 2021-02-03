package arrayStack

import (
	"bytes"
	"fmt"
	"github.com/algo/DataStructureByGo/02Stacks-and-Queues/01Stack/array"
)

type ArrayStack struct {
	*array.MyArray
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		array.NewArray(capacity),
	}
}

// 获取栈中元素数目
func (as *ArrayStack) GetStackSize() int {
	return as.GetSize()
}

// 判断栈是否为空
func (as *ArrayStack) IsStackEmpty() bool {
	return as.IsEmpty()
}

// 压栈, 底层数组能够自动扩缩容
func (as *ArrayStack) Push(e interface{}) {
	as.AddLast(e)
}

// 出栈并返回栈顶元素
func (as *ArrayStack) Pop() interface{} {
	return as.DeleteLast()
}

// 返回栈顶元素
func (as *ArrayStack) Peek() interface{} {
	return as.Get(as.GetSize() - 1)
}

// 将栈以字符串的形式打印出来
func (as *ArrayStack) ArrayStackToString() string {
	var buf bytes.Buffer
	buf.WriteString("Stack:")
	if as.GetSize() == 0 {
		buf.WriteString("[] top")
	} else {
		buf.WriteString("[")
		for i := 0; i < as.GetSize(); i++ {
			buf.WriteString(fmt.Sprintf("%v", as.Get(i)))
			if i != as.GetSize()-1 {
				buf.WriteString(" ,")
			}
		}
		buf.WriteString("] top")
	}
	return buf.String()
}
