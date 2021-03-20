package main

import (
	"01MaxHeap/maxHeap"
	"01MaxHeap/util"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func testHeap(testData []interface{}, isHeapify bool) time.Duration {
	startTime := time.Now()

	var heap *maxHeap.MaxHeap
	if isHeapify {
		heap = maxHeap.NewMaxHeapFromArr(testData)
	} else {
		heap = maxHeap.NewMaxHeap()
		for _, v := range testData {
			heap.Add(v)
		}
	}

	arr := make([]interface{}, len(testData))
	for i := 0; i < len(testData); i++ {
		arr[i] = heap.ExtractMax()
	}
	for i := 1; i < len(testData); i++ {
		if util.Compare(arr[i-1], arr[i]) < 0 {
			panic("Error")
		}
	}

	fmt.Println("test MaxHeap completed")

	return time.Now().Sub(startTime)
}

func main() {
	n := 1000000

	var testData []interface{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		testData = append(testData, rand.Intn(math.MaxInt32))
	}

	time1 := testHeap(testData, false)
	fmt.Println("Without heapify: ", time1)

	time2 := testHeap(testData, true)
	fmt.Println("With heapify: ", time2)
}
