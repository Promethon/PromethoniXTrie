package main

type AccountData struct {
	Address  [20]byte
	BlockNum uint64
	Value    int64
	Data     []byte
}
