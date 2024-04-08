package benchmark

type EmptyTreeStruct struct {
}

func (tree EmptyTreeStruct) Get(string) (TestData, error) {
	return TestData{}, nil
}

func (tree EmptyTreeStruct) Add(string, uint64, []byte) error {
	return nil
}

func (tree EmptyTreeStruct) Delete(string) error {
	return nil
}

func (tree EmptyTreeStruct) UpdateBlockHeight(uint64) error {
	return nil
}
