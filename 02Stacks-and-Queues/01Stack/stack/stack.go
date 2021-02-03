package stack

type Stack interface {
	GetStackSize() int
	IsStackEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	ArrayStackToString() string
}
