package main

import (
	"PromethoniXTrie/benchmark"
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

func LoadDataset(tree benchmark.TestStruct) error {
	file, err := os.Open(DatasetName)
	if err != nil {
		return err
	}
	defer file.Close()

	dec := gob.NewDecoder(file)

	for i := 0; i < SampleSize; i++ {
		var s AccountData
		err = dec.Decode(&s)

		if s.BlockNum != 0 {
			tree.UpdateBlockHeight(s.BlockNum)
			tree.Add(string(s.Address[:]), uint64(s.Value), s.Data)
		}
	}
	return err
}

func main() {
	fmt.Println("Generating Random Dataset...")
	generateRandomDataset()

	// dummy test
	LoadDataset(benchmark.EmptyTreeStruct{})

	fmt.Println("Starting...")
	var startTime int64

	promethon, _ := benchmark.NewPromethoniXTrieImpl(nil, false)
	startTime = time.Now().UnixNano()
	LoadDataset(promethon)
	fmt.Printf("Promethon:\t%d\n", time.Now().UnixNano()-startTime)
}
