package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// 插入一个前缀
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// 判断字符串是否匹配任意前缀
func (t *Trie) HasAnyPrefix(s string) bool {
	node := t.root
	for _, ch := range s {
		node = node.children[ch]
		if node == nil {
			return false
		}
		if node.isEnd {
			return true
		}
	}
	return false
}

// 返回匹配到的前缀
func (t *Trie) MatchPrefix(s string) (string, bool) {
	node := t.root
	runes := []rune(s)
	for i, ch := range runes {
		node = node.children[ch]
		if node == nil {
			return "", false
		}
		if node.isEnd {
			return string(runes[:i+1]), true
		}
	}
	return "", false
}
