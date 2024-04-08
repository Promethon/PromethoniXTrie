package benchmark

import (
	"PromethoniXTrie"
	"PromethoniXTrie/redblacktree"
	"unsafe"
)

type TrieTreeImpl struct {
	Trie        *PromethoniXTrie.PromethoniXTrie
	Tree        *redblacktree.Tree
	BlockHeight uint64
}

func NewPromethoniXTrieImpl(
	rootHash PromethoniXTrie.Hash,
	isActionLogEnabled bool,
) (*TrieTreeImpl, error) {
	trie, err := PromethoniXTrie.NewPromethoniXTrie("db", rootHash, isActionLogEnabled)
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

func comparator(a, b interface{}) int {
	a1, ok := a.(TestData)
	if !ok {
		return 0
	}
	b2, ok := b.(TestData)
	if !ok {
		return 0
	}

	aAsserted := a1.Value + a1.LastModified
	bAsserted := b2.Value + b2.LastModified
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func (t *TrieTreeImpl) Add(key string, value uint64, data []byte) error {
	gottenData, err := t.Trie.Get([]byte(key))
	if err == nil {
		// EncodedData already exist
		rbKey, ok := ((*redblacktree.Node)(unsafe.Pointer(byteSliceToPtr(gottenData))).Key).(TestData)
		newData := data
		if ok {
			newData = append(rbKey.Data, data...)
			t.Tree.Remove(rbKey)
		}

		t.Trie.Delete([]byte(key))
		RBNode := t.Tree.Put(TestData{Value: value, LastModified: t.BlockHeight, Data: newData}, key)
		_, err = t.Trie.Put([]byte(key), ptrToByteSlice(uintptr(unsafe.Pointer(RBNode))))
	} else {
		// EncodedData does not exist
		RBNode := t.Tree.Put(TestData{Value: value, LastModified: t.BlockHeight, Data: data}, key)
		_, err = t.Trie.Put([]byte(key), ptrToByteSlice(uintptr(unsafe.Pointer(RBNode))))
	}
	return err
}

func (t *TrieTreeImpl) Get(key string) (TestData, error) {
	gottenData, err := t.Trie.Get([]byte(key))
	if err != nil {
		return TestData{}, err
	}
	fs2 := (*redblacktree.Node)(unsafe.Pointer(byteSliceToPtr(gottenData)))
	return fs2.Key.(TestData), nil
}

func (t *TrieTreeImpl) Delete(key string) error {
	gottenData, err := t.Trie.Get([]byte(key))
	if err == nil {
		t.Tree.Remove2((*redblacktree.Node)(unsafe.Pointer(byteSliceToPtr(gottenData))))
		_, err = t.Trie.Delete([]byte(key))
	}
	return err
}

func (t *TrieTreeImpl) UpdateBlockHeight(newHeight uint64) error {
	t.BlockHeight = newHeight
	for {
		if t.Tree.Left() == nil {
			return nil
		}
		first := t.Tree.Left().Key.(TestData).Value
		second := t.BlockHeight - t.Tree.Left().Key.(TestData).LastModified
		if first <= second {
			if err := t.Delete(t.Tree.Left().Value.(string)); err != nil {
				return err
			}
		} else {
			return nil
		}
	}
}
