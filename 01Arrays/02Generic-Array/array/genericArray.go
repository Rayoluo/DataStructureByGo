// 实现支持泛型的数组
package array

import (
	"bytes"
	"fmt"
	"reflect"
)

type MyArray struct {
	arr  []interface{}
	size int
}

// 创建一个新的数组, 切片长度设置为capacity, 相当于定长数组
func NewArray(capacity int) *MyArray {
	return &MyArray{
		arr:  make([]interface{}, capacity),
		size: 0,
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
func (m *MyArray) AddLast(elem interface{}) {
	m.AddIndex(elem, m.size)
}

// 在数组第index个位置添加一个元素, index 0 ~ size
func (m *MyArray) AddIndex(elem interface{}, index int) {
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

// 在数组开始位置添加一个元素
func (m *MyArray) AddFirst(elem interface{}) {
	m.AddIndex(elem, 0)
}

// 获取数组指定位置的元素，index范围为[0,size)
func (m *MyArray) Get(index int) interface{} {
	if index < 0 || index >= m.size {
		panic("Illegal index value!")
	}
	return m.arr[index]
}

// 修改数组指定位置的元素
func (m *MyArray) Set(index int, value interface{}) {
	if index < 0 || index >= m.size {
		panic("Illegal index value!")
	}
	m.arr[index] = value
}

// 查找数组中是否包含指定的元素 待修改
func (m *MyArray) Contains(e interface{}) bool {
	for i := 0; i < m.size; i++ {
		if reflect.DeepEqual(m.arr[i], e) {
			return true
		}
	}
	return false
}

// 查找数组中指定元素的索引，如果不存在元素则返回-1，待修改
func (m *MyArray) Find(e interface{}) int {
	for i := 0; i < m.size; i++ {
		if reflect.DeepEqual(m.arr[i], e) {
			return i
		}
	}
	return -1
}

// 删除数组中指定位置的元素, 位置只能是0 - (size-1)之间的元素, 返回删除的元素
func (m *MyArray) DeleteIndex(index int) (ret interface{}) {
	if index < 0 || index >= m.size {
		panic("DeleteIndex Fail: Illegal index value!")
	}
	ret = m.arr[index]
	for i := index + 1; i < m.size; i++ {
		m.arr[i-1] = m.arr[i]
	}
	m.size--
	return
}

// 删除数组首个元素
func (m *MyArray) DeleteFirst() (ret interface{}) {
	ret = m.DeleteIndex(0)
	return
}

// 删除数组末尾元素
func (m *MyArray) DeleteLast() (ret interface{}) {
	ret = m.DeleteIndex(m.size - 1)
	return
}

// 将数组以字符串的形式打印
func (m *MyArray) ArrayToString() string {
	var buf bytes.Buffer
	if size := m.GetSize(); size == 0 {
		buf.WriteString("[]")
	} else {
		buf.WriteString("[")
		for i := 0; i < size; i++ {
			if i == size-1 {
				buf.WriteString(fmt.Sprintf("%v", m.arr[i]) + "]")
			} else {
				buf.WriteString(fmt.Sprintf("%v", m.arr[i]) + " ")
			}
		}
	}
	return buf.String()
}
