# 剑指offer栈和队列类型题解

## 面试题09 用两个栈实现队列

题目链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

双栈实现队列：

**成员变量**

+ 维护两个栈 stack1 和 stack2，其中 stack1 支持插入操作，stack2 支持删除操作

**构造方法**

+ 初始化 stack1 和 stack2 为空

**插入元素**

插入元素对应方法 appendTail

+ stack1 直接插入元素

**删除元素**

删除元素对应方法 deleteHead

+ 如果 stack2 为空，则将 stack1 里的所有元素弹出插入到 stack2 里
+ 如果 stack2 仍为空，则返回 -1，否则从 stack2 弹出一个元素并返回

![示意图](https://assets.leetcode-cn.com/solution-static/jianzhi_09/jianzhi_9.gif)

实现：

```go
type CQueue struct {
	stack1     []int
	stack2     []int
	top1, top2 int
}

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 10),
		stack2: make([]int, 10),
	}
}

func (this *CQueue) AppendTail(value int) {
	// 数组满了,则扩容
	if this.top1 == len(this.stack1) {
		this.stack1 = resize(this.stack1, this.top1, len(this.stack1)*2)
	}
	this.stack1[this.top1] = value
	this.top1++
}

func resize(arr []int, top int, newCapacity int) []int {
	newArr := make([]int, newCapacity)
	for i := 0; i < top; i++ {
		newArr[i] = arr[i]
	}
	return newArr
}

func (this *CQueue) DeleteHead() int {
	// 队列为空的情况
	if this.top1 == 0 && this.top2 == 0 {
		return -1
	}
	if this.top1 != 0 && this.top2 == 0 {
		for this.top1 != 0 {
			value := this.stack1[this.top1-1]
			this.top1--
			// 数组满了则扩容
			if this.top2 == len(this.stack2) {
				this.stack2 = resize(this.stack2, this.top2, len(this.stack2)*2)
			}
			this.stack2[this.top2] = value
			this.top2++
		}
	}
	ret := this.stack2[this.top2-1]
	this.top2--
	if this.top1 == len(this.stack1)/4 && len(this.stack1)/2 != 0 {
		this.stack1 = resize(this.stack1, this.top1, len(this.stack1)/2)
	}
	if this.top2 == len(this.stack2)/4 && len(this.stack2)/2 != 0 {
		this.stack2 = resize(this.stack2, this.top2, len(this.stack2)/2)
	}
	return ret
}
```

这里的实现稍微有些复杂，原因是引入了`top1`和`top2`指针，将切片看作数组处理，手动对数组进行扩缩容。没有充分利用到切片的自动扩容的特性，下面的实现去掉`top1`和`top2`指针，使用`len(this.stack1)`表示栈中元素数目，`this.stack1 = this.stack1[:len(this.stack1)-1]`表示从`stack1`的栈顶删除一个元素。

```go
type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	// 队列为空的情况
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return -1
	}
	if len(this.stack1) != 0 && len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			value := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, value)
		}
	}
	ret := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return ret
}
```

## 面试题30 包含min函数的栈

题目链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

题目描述：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的`min`函数在该栈中，调用`min`、`push`及`pop`的时间复杂度都是`O(1)`

这道题的想法首先是能不能用一个变量记录当前状态的栈的最小元素，但是问题在于当从栈中弹出一个元素之后`min`返回的应该是栈的次小元素，可是它并没有被记录。那么这里该如何记录这些信息呢？一个思路是使用小顶堆来记录这些信息，另一个思路是使用一个辅助栈，**记录每次压栈操作的最小元素**。这里采取第二种思路来实现。

实现：

```go
type MinStack struct {
	stack1 []int // 数据栈
	stack2 []int // 辅助栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack2) == 0 {
		this.stack2 = append(this.stack2, x)
	} else {
		minVal := this.stack2[len(this.stack2)-1]
		if x < minVal {
			this.stack2 = append(this.stack2, x)
		} else {
			this.stack2 = append(this.stack2, minVal)
		}
	}
	this.stack1 = append(this.stack1, x)
}

func (this *MinStack) Pop() {
	if len(this.stack1) == 0 {
		panic("Error! Empty Stack!")
	}
	this.stack1 = this.stack1[:len(this.stack1)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack1) == 0 {
		panic("Error! Empty Stack!")
	}
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack2) == 0 {
		panic("Error! Empty Stack!")
	}
	return this.stack2[len(this.stack2)-1]
}
```

## 面试题31 栈的压入、弹出序列

题目链接：https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

题目描述：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

这道题我的思路是纯模拟。依次读弹出序列，分三种情况：

+ 如果弹出序列当前元素未压入栈，将压入序列中该元素之前的未访问到的所有元素压入栈中；
+ 如果弹出序列当前元素已经压入栈中但不是栈顶元素，返回`false`；
+ 如果弹出序列当前元素已经压入栈中且为栈顶元素，弹出该栈顶元素，继续访问下一个弹出序列元素；

实现：

```go
func validateStackSequences(pushed []int, popped []int) bool {
	mp := make(map[int]bool) // 标记元素是否已经压入栈
	stack := make([]int, 0)
	visitIndex := 0
	for i := 0; i < len(popped); i++ {
		// 如果该元素没有被压入栈，则压入序列中该元素之前的所有元素都将被压入栈中并被标记
		if _, ok := mp[popped[i]]; !ok {
			for j := visitIndex; j < len(pushed) && pushed[j] != popped[i]; j++ {
				stack = append(stack, pushed[j])
				mp[pushed[j]] = true
				visitIndex++
			}
			mp[popped[i]] = true
			visitIndex++
			continue
		} else {
			// 如果元素已被压入栈中，但不是栈顶元素，则返回false
			if topElem := stack[len(stack)-1]; topElem != popped[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return true
}
```

这个算法的空间复杂度很拉胯，提交仅击败7%。如何进一步优化空间复杂度呢，实际上标记映射并不是必须的。只要将弹出序列的所有元素都访问了一遍说明就是合理的。实现：

```go
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	visitIndex := 0
	for i := 0; i < len(popped); i++ {
		if len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
		} else {
			for j := visitIndex; j < len(pushed) && pushed[j] != popped[i]; j++ {
				stack = append(stack, pushed[j])
				visitIndex++
			}
			if visitIndex == len(pushed) {
				return false
			}
			visitIndex++
		}
	}
	return true
}
```

<img src="https://gitee.com/oluoluo/typoraImage/raw/master/img/image-20210204201346372.png" alt="image-20210204201346372" style="zoom:67%;" />

## 面试题59 队列的最大值

题目链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

题目描述：定义一个队列并实现函数`max_value`得到队列里的最大值，要求函数`max_value`, `push_back`和`pop_front`的均摊时间复杂度都是`O(1)`；若队列为空，`pop_front`和`max_value`需要返回-1

这道题和前面的包含`min`函数的栈很像，即包含`max`函数的队列。思维历程也是相似的，首先如果使用一个变量记录当前的最大值，如果一个元素入队了就将该元素与记录变量进行比较然后决定是否更新该变量，但是当有一个元素出队了并且该元素刚好是队列的最大元素，那么此时队列中次大的元素我们并不知道。咋整呢？没想出来，看了一下题解，是使用一个双端队列，思路rt，很巧妙:

![示意图](https://pic.leetcode-cn.com/9d038fc9bca6db656f81853d49caccae358a5630589df304fc24d8999777df98-fig3.gif)

```go
type MaxQueue struct {
	dataQueue []int // 普通队列
	deque     []int // 双端队列
}

func Constructor() MaxQueue {
	return MaxQueue{
		dataQueue: make([]int, 0),
		deque:     make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.dataQueue = append(this.dataQueue, value)
	i := len(this.deque) - 1
	for i >= 0 {
		if this.deque[i] > value {
			break
		}
		i--
	}
	this.deque = append(this.deque[:i+1], value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.dataQueue) == 0 {
		return -1
	}
	ret := this.dataQueue[0]
	if this.dataQueue[0] == this.deque[0] {
		this.deque = this.deque[1:]
	}
	this.dataQueue = this.dataQueue[1:]
	return ret
}
```





