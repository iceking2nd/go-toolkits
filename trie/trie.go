package trie

import "sync"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
	mu   sync.RWMutex
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// 插入前缀（写锁）
func (t *Trie) Insert(word string) {
	t.mu.Lock()
	defer t.mu.Unlock()

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

// 判断是否匹配任意前缀（读锁）
func (t *Trie) HasAnyPrefix(s string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

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

// 返回匹配到的前缀（读锁）
func (t *Trie) MatchPrefix(s string) (string, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

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
