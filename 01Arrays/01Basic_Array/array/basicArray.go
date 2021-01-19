package array

import (
	"bytes"
	"strconv"
)

type MyArray struct {
	arr  []int
	size int
}

// 创建一个新的数组, 切片长度设置为capacity, 相当于定长数组
func NewArray(capacity int) *MyArray {
	return &MyArray{
		arr: make([]int, capacity),
	}
}

// 获取数组容量
func (m *MyArray) GetCapacity() int {
	return len(m.arr)
}

// 获取数组中元素的个数
func (m *MyArray) GetSize() int {
	return m.size
}

// 判断数组是否为空
func (m *MyArray) IsEmpty() bool {
	return m.size == 0
}

// 向数组末尾位置添加一个元素
func (m *MyArray) AddLast(elem int) {
	// 如果数组满了就抛出异常
	if m.size == m.GetCapacity() {
		panic("AddLast fail: full array!")
	}
	m.arr[m.size] = elem
	m.size++
}

// 在数组第index个位置添加一个元素, index 0 ~ size
func (m *MyArray) AddIndex(elem, index int) {
	if m.GetCapacity() == m.size {
		panic("AddIndex fail: full array!")
	}
	if index < 0 || index > m.size {
		panic("Illegal parameter index!")
	}
	for i := m.size - 1; i >= index; i-- {
		m.arr[i+1] = m.arr[i]
	}
	m.arr[index] = elem
	m.size++
}

// 将数组以字符串的形式打印
func (m *MyArray) ArrayToString() string {
	var buf bytes.Buffer
	if size := m.GetSize(); size == 0 {
		buf.WriteString("[]")
	} else {
		buf.WriteString("[")
		for i, v := range m.arr[:size] {
			if i == size-1 {
				buf.WriteString(strconv.Itoa(v) + "]")
			} else {
				buf.WriteString(strconv.Itoa(v) + " ")
			}
		}
	}
	return buf.String()
}
