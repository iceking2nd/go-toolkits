package trie

import (
	"fmt"
	"sync"
	"testing"
)

func TestHasAnyPrefix(t *testing.T) {
	trie := NewTrie()

	prefixes := []string{"http://", "https://", "ftp://"}
	var wg sync.WaitGroup

	// 并发插入
	for _, p := range prefixes {
		wg.Add(1)
		go func(prefix string) {
			defer wg.Done()
			trie.Insert(prefix)
		}(p)
	}
	wg.Wait()

	// 并发读取
	testStr := "https://example.com"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if match, ok := trie.MatchPrefix(testStr); ok {
				fmt.Println("匹配:", match)
			}
		}()
	}
	wg.Wait()
}
