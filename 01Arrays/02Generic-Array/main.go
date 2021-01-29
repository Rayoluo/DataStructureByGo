package main

import (
	"fmt"
	"github.com/algo/DataStructureByGo/01Arrays/02Generic-Array/array"
)

type student struct {
	name  string
	score int
}

func main() {
	arr := array.NewArray(10)
	arr.AddLast(student{
		name:  "Wanghuining",
		score: 99,
	})
	arr.AddLast(student{
		name:  "Liuyulan",
		score: 100,
	})
	arr.AddFirst(student{
		name:  "Huyuewei",
		score: 96,
	})
	fmt.Println(arr.ArrayToString())
	fmt.Println("Contains student Huyuewei with score 96?", arr.Contains(student{
		name:  "Huyuewei",
		score: 96,
	}))
	fmt.Println("Find location of student wanghuining:", arr.Find(student{
		name:  "Wanghuining",
		score: 99,
	}))
}
