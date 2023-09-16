package PromethoniXTrie

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestWriteAndReadNodeType(t *testing.T) {
	var buf bytes.Buffer
	var testValue NodeType = 1

	err := writeNodeType(&buf, testValue)
	if err != nil {
		t.Fatalf("writeNodeType() error = %v", err)
	}

	value, err := readNodeType(&buf)
	if err != nil {
		t.Fatalf("readNodeType() error = %v", err)
	}

	if value != testValue {
		t.Errorf("Expected %v, but got %v", testValue, value)
	}
}

func TestWriteAndReadInt32(t *testing.T) {
	var buf bytes.Buffer
	var testValue int32 = 1234

	err := writeInt32(&buf, testValue)
	if err != nil {
		t.Fatalf("writeInt32() error = %v", err)
	}

	value, err := readInt32(&buf)
	if err != nil {
		t.Fatalf("readInt32() error = %v", err)
	}

	if value != testValue {
		t.Errorf("Expected %v, but got %v", testValue, value)
	}
}

func TestReadBytes(t *testing.T) {
	var buf bytes.Buffer
	testBytes := []byte{1, 2, 3, 4, 5}

	buf.Write(testBytes)

	readBytes, err := readBytes(&buf, int32(len(testBytes)))
	if err != nil {
		t.Fatalf("readBytes() error = %v", err)
	}

	for i, b := range readBytes {
		if b != testBytes[i] {
			t.Errorf("Expected %v, but got %v", testBytes[i], b)
		}
	}
}

func TestSha3Hash(t *testing.T) {
	hash, _ := sha3Hash([]byte("test data"))
	if hex.EncodeToString(hash) != "fc88e0ac33ff105e376f4ece95fb06925d5ab20080dbe3aede7dd47e45dfd931" {
		t.Error("expected fc88e0ac33ff105e376f4ece95fb06925d5ab20080dbe3aede7dd47e45dfd931 but got ", hex.EncodeToString(hash))
	}
}

func TestKeyToRouteAndRouteToKey(t *testing.T) {
	testHash := Hash{0xa1, 0xf2}
	route := keyToRoute(testHash)
	key := routeToKey(route)

	if !bytes.Equal(key, testHash) {
		t.Errorf("Expected %v, but got %v", testHash, key)
	}
}

func TestPrefixLen(t *testing.T) {
	a := []byte{1, 2, 3, 4, 5}
	b := []byte{1, 2, 3, 0, 0}
	expectedLen := 3

	length := prefixLen(a, b)
	if length != expectedLen {
		t.Errorf("Expected %v, but got %v", expectedLen, length)
	}
}

func TestPtrToByteSliceAndByteSliceToPtr(t *testing.T) {
	testPtr := uintptr(123456)
	b := ptrToByteSlice(testPtr)
	ptr := byteSliceToPtr(b)

	if ptr != testPtr {
		t.Errorf("Expected %v, but got %v", testPtr, ptr)
	}
}
