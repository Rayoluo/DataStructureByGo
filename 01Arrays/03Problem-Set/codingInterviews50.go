package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abaccdeff"
	fmt.Printf("%c\n", firstUniqChar1(s))
}

type pair struct {
	ch         byte
	cnt, order int
}

// s只包含小写字母, 时间复杂度是O(nlogn)并非最佳
func firstUniqChar(s string) byte {
	hashtable := make([]pair, 26)
	oid := 1
	for i := 0; i < len(s); i++ {
		if hashtable[s[i]-'a'].cnt == 0 {
			hashtable[s[i]-'a'].order = oid
			oid++
			hashtable[s[i]-'a'].ch = s[i]
		}
		hashtable[s[i]-'a'].cnt++
	}
	sort.Slice(hashtable, func(i, j int) bool {
		return hashtable[i].order < hashtable[j].order
	})
	for i := 0; i < 26; i++ {
		if hashtable[i].cnt == 1 {
			return hashtable[i].ch
		}
	}
	return ' '
}

// 第二种解法，访问两遍字符串，不使用map执行时间更短
func firstUniqChar1(s string) byte {
	// mp := make(map[byte]int)
	// for i := 0; i < len(s); i++ {
	// 	mp[s[i]] = mp[s[i]] + 1
	// }
	// for i := 0; i < len(s); i++ {
	// 	if mp[s[i]] == 1 {
	// 		return s[i]
	// 	}
	// }
	var hashtable [26]int
	for i := 0; i < len(s); i++ {
		hashtable[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if hashtable[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

//
