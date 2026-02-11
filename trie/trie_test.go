package trie

import (
	"fmt"
	"testing"
)

func TestHasAnyPrefix(t *testing.T) {
	prefixes := []string{"http://", "https://", "ftp://"}

	trie := NewTrie()
	for _, p := range prefixes {
		trie.Insert(p)
	}

	testStr := "https://example.com"

	fmt.Println(trie.HasAnyPrefix(testStr))

	if match, ok := trie.MatchPrefix(testStr); ok {
		fmt.Println("匹配到前缀:", match)
	}
}
