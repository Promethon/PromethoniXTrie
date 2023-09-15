package PromethoniXTrie

import (
	"PromethoniXTrie/redblacktree"
	"unsafe"
)

type TrieTreeImpl struct {
	Trie *PromethoniXTrie
	Tree *redblacktree.Tree
}

func NewTrieTreeImpl(
	rootHash Hash,
	isActionLogEnabled bool,
) (*TrieTreeImpl, error) {
	trie, err := NewPromethoniXTrie(rootHash, isActionLogEnabled)
	if err != nil {
		return nil, err
	}

	trieTree := &TrieTreeImpl{
		Trie: trie,
		Tree: &redblacktree.Tree{
			Comparator: comparator,
		},
	}
	return trieTree, nil
}

type Comparable struct {
	Value        uint64
	LastModified uint64
}

func comparator(a, b interface{}) int {
	aAsserted := a.(Comparable).Value - (blockHeight - a.(Comparable).LastModified)
	bAsserted := b.(Comparable).Value - (blockHeight - b.(Comparable).LastModified)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func (t *TrieTreeImpl) Add(key string, value uint64) error {
	gottenData, err := t.Trie.Get([]byte(key))
	if err == nil {
		// EncodedData already exist
		t.Tree.Remove((*redblacktree.Node)(unsafe.Pointer(byteSliceToPtr(gottenData))).Key)
		t.Trie.Delete([]byte(key))
		RBNode := t.Tree.Put(Comparable{Value: value, LastModified: blockHeight}, key)
		_, err = t.Trie.Put([]byte(key), ptrToByteSlice(uintptr(unsafe.Pointer(RBNode))))
	} else {
		// EncodedData does not exist
		RBNode := t.Tree.Put(Comparable{Value: value, LastModified: blockHeight}, key)
		_, err = t.Trie.Put([]byte(key), ptrToByteSlice(uintptr(unsafe.Pointer(RBNode))))
	}
	return err
}

func (t *TrieTreeImpl) Get(key string) (uint64, uint64, error) {
	gottenData, err := t.Trie.Get([]byte(key))
	if err != nil {
		return 0, 0, err
	}
	fs2 := (*redblacktree.Node)(unsafe.Pointer(byteSliceToPtr(gottenData)))
	return fs2.Key.(Comparable).Value, fs2.Key.(Comparable).LastModified, nil
}

func (t *TrieTreeImpl) Delete(key string) error {
	gottenData, err := t.Trie.Get([]byte(key))
	if err == nil {
		t.Tree.Remove((*redblacktree.Node)(unsafe.Pointer(byteSliceToPtr(gottenData))).Key)
		_, err = t.Trie.Delete([]byte(key))
	}
	return err
}

var blockHeight uint64 = 909

func (t *TrieTreeImpl) ChangeBlockHeight(newHeight uint64) error {
	blockHeight = newHeight
	for {
		first := t.Tree.Left().Key.(Comparable).Value
		second := blockHeight - t.Tree.Left().Key.(Comparable).LastModified
		if first <= second {
			if err := t.Delete(t.Tree.Left().Value.(string)); err != nil {
				return err
			}
		} else {
			return nil
		}
	}
}
