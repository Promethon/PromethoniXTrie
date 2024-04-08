package main

import (
	"crypto/rand"
	"encoding/gob"
	"log"
	r2 "math/rand"
	"os"
)

const DatasetName = "dataset"
const AddressCount = 10
const DataSize = 1000
const SampleSize = 10000
const Epoch = SampleSize / 2

func appendToFile(filename string, data AccountData) error {
	// Open the file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create(DatasetName)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()

	// Create a new gob encoder
	enc := gob.NewEncoder(file)

	// Encode the data
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

var blocknum uint64
var addresses [][20]byte = nil

func randgenerator() {
	data := make([]byte, r2.Intn(DataSize))
	rand.Read(data[:])
	blocknum += 1

	newData := AccountData{
		BlockNum: blocknum,
		Address:  addresses[r2.Intn(AddressCount)],
		Value:    r2.Int63n(9223372036854775807),
		Data:     data,
	}

	appendToFile(DatasetName, newData)
}

func generateRandomDataset() {
	os.Remove(DatasetName)

	addresses = make([][20]byte, AddressCount)
	for i := 0; i < AddressCount; i++ {
		var address [20]byte
		rand.Read(address[:])
		addresses[i] = address
	}

	for i := 0; i < SampleSize; i++ {
		randgenerator()
	}
}
