package main

import (
	"PromethoniXTrie"
	"fmt"
)

func main() {

	t, _ := PromethoniXTrie.NewTrieTreeImpl(nil, true)
	fmt.Printf("trie2:root1: %x \n", t.Trie.RootHash())
	t.Add("0xeaaaffee24f3ce7b942e7016e37ea2899a3004df", 2526372489)
	t.ChangeBlockHeight(1990)
	t.Add("0xeaac8635e9e62ff2e9650881fa37708bc06bdea4", 15372334787)
	t.ChangeBlockHeight(899886)
	fmt.Printf("trie root before adding 0x42cc07ded70beed62c2dffd75724619c34ced8cc: %x \n", t.Trie.RootHash())
	t.Add("0x42cc07ded70beed62c2dffd75724619c34ced8cc", 112)
	fmt.Printf("trie root after adding 0x42cc07ded70beed62c2dffd75724619c34ced8cc: %x \n", t.Trie.RootHash())
	fmt.Println(t.Tree)
	t.ChangeBlockHeight(8998865)
	fmt.Printf("trie root after deleting 0x42cc07ded70beed62c2dffd75724619c34ced8cc: %x \n", t.Trie.RootHash())
	fmt.Println(t.Tree)
	fmt.Println(t.Trie.ActionLogEntries)
}
