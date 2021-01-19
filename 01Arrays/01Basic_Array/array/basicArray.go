package array

type MyArray struct {
	arr []int
}

// 创建一个新的数组
func NewArray(capacity int) *MyArray {
	return &MyArray{
		arr:make([]int, 0, capacity),
	}
}

// 获取数组容量
func (m *MyArray) GetCapacity() int {
	return cap(m.arr)
}

// 获取数组中元素的个数
func (m *MyArray) GetSize() int {
	return len(m.arr)
}

// 判断数组是否为空
func (m *MyArray) IsEmpty() bool {
	return len(m.arr) == 0
}