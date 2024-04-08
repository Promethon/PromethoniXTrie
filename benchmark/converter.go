package benchmark

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

type TestData struct {
	Value        uint64
	LastModified uint64
	Data         []byte
}

type TestStruct interface {
	Get(string) (TestData, error)
	Add(string, uint64, []byte) error
	Delete(string) error
	UpdateBlockHeight(uint64) error
}

func encode(node TestData) ([]byte, error) {
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.LittleEndian, node.Value)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buf, binary.LittleEndian, node.LastModified)
	if err != nil {
		return nil, err
	}

	var l uint32
	l = uint32(len(node.Data))
	err = binary.Write(&buf, binary.LittleEndian, l)
	if err != nil {
		return nil, err
	}
	if err == nil && l > 0 {
		_, err = buf.Write(node.Data)
	}

	return buf.Bytes(), err
}

func decode(raw []byte) (TestData, error) {
	buf := bytes.NewReader(raw)
	var value uint64
	var lastModified uint64
	var length uint32
	out := TestData{}

	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		return out, err
	}
	out.Value = value

	err = binary.Read(buf, binary.LittleEndian, &lastModified)
	if err != nil {
		return out, err
	}
	out.LastModified = lastModified

	err = binary.Read(buf, binary.LittleEndian, &length)
	if err != nil {
		return TestData{}, err
	}

	if length > 0 {
		data := make([]byte, length)
		_, err = buf.Read(data)
		out.Data = data
	}

	return out, err
}

func ptrToByteSlice(u uintptr) []byte {
	size := unsafe.Sizeof(u)
	b := make([]byte, size)
	switch size {
	case 4:
		binary.LittleEndian.PutUint32(b, uint32(u))
	case 8:
		binary.LittleEndian.PutUint64(b, uint64(u))
	}
	return b
}

func byteSliceToPtr(b []byte) uintptr {
	switch len(b) {
	case 4:
		return uintptr(binary.LittleEndian.Uint32(b))
	case 8:
		return uintptr(binary.LittleEndian.Uint64(b))
	}
	return 0
}
