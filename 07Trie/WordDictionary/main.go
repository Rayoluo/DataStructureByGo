package main

// Leetcode 211 添加与搜索单词 数据结构设计 word中可能包含.通配符

type Node struct {
	isWord bool
	next   map[string]*Node
}

type WordDictionary struct {
	root *Node
	size int
}

func generateNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: generateNode(),
	}
}

// 向Trie中添加一个单词word
func (wd *WordDictionary) AddWord(word string) {
	cur := wd.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = generateNode()
		}
		cur = cur.next[c]
	}
	if cur.isWord == false {
		cur.isWord = true
		wd.size++
	}
}

// 查询单词word是否在trie中, 可以匹配.操作符
func (wd *WordDictionary) Search(word string) bool {
	return wd.subSearch(wd.root, word, 0)
}

func (wd *WordDictionary) subSearch(root *Node, word string, index int) bool {
	if index >= len([]rune(word)) {
		return root.isWord
	}
	cur := root
	resword := ([]rune(word))[index:]
	for _, w := range resword {
		c := string(w)
		if c != "." {
			if cur.next[c] == nil {
				return false
			}
			cur = cur.next[c]
			index++
		} else {
			for _, val := range cur.next {
				if wd.subSearch(val, word, index+1) {
					return true
				}
			}
			return false
		}
	}
	return cur.isWord
}
