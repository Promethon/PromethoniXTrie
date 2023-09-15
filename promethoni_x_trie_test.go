package PromethoniXTrie

import (
	"testing"
)

func TestEmptyTrie(t *testing.T) {
	trie, _ := NewPromethoniXTrie(nil, false)
	res := trie.RootHash()
	if res != nil {
		t.Errorf("expected nil but got %x", res)
	}
}

func TestIsEmptyTrie(t *testing.T) {
	trie, _ := NewPromethoniXTrie(nil, false)
	if !trie.IsEmpty() {
		t.Errorf("expected trie to be empty but was not")
	}
}

func TestNullKey(t *testing.T) {
	trie, _ := NewPromethoniXTrie(nil, false)
	_, err := trie.Put(nil, []byte("test value"))
	if err == nil {
		t.Errorf("expected error but didn't get")
	}
}

func TestNullValue(t *testing.T) {
	trie, _ := NewPromethoniXTrie(nil, false)
	_, err := trie.Put([]byte("test key"), nil)
	if err == nil {
		t.Errorf("expected error but didn't get")
	}
}
