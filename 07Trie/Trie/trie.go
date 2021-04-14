package Trie

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
	size int
}

func generateNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

func New() *Trie {
	return &Trie{
		root: &Node{
			next: make(map[string]*Node),
		},
	}
}

// 获得Trie中存储的单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// 向Trie中添加一个单词word
func (t *Trie) Add(word string) {
	cur := t.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = generateNode()
		}
		cur = cur.next[c]
	}
	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}

// 查询单词word是否在trie中
func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}

// 前缀查询: 判断某个单词以这个字符串开始
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return true // 只要找到这个前缀，就可以认为这个前缀的单词是存在的，不需要判断前缀到达的位置是否是单词
}
