package set

type Set interface {
	Add(interface{})
	IsEmpty() bool
	Contains(interface{}) bool
	GetSize() int
	Remove(interface{})
}
