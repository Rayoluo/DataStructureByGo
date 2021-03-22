package main

import (
	"fmt"
)

func main() {
	str := "google"
	for i := 0; i < len(str); i++ {
		Insert(byte(str[i]))
		fmt.Printf("%c\n", FirstAppearingOnce())
		// fmt.Println(FirstAppearingOnce())
	}
}

var (
	countMp map[byte]int
	orderMp map[int]byte
	orderId int
)

func Insert(ch byte) {
	if countMp == nil || orderMp == nil {
		countMp = make(map[byte]int)
		orderMp = make(map[int]byte)
	}
	orderMp[orderId] = ch
	orderId++
	if val, ok := countMp[ch]; !ok {
		countMp[ch] = 1
	} else {
		countMp[ch] = val + 1
	}
}

func FirstAppearingOnce() byte {
	for i := 0; i < len(orderMp); i++ {
		if countMp[orderMp[i]] == 1 {
			return orderMp[i]
		}
	}
	return byte('#')
}

// 另一种方法高效一些：即用一个哈希表存放字符出现的次数，用一个队列放字符流出现的顺序，如果是出现多次的
// 字符就不push到队列中
