package gtk

type SelectionData interface {
	GetData() []byte
	GetLength() int
} // end of SelectionData

func AssertSelectionData(_ SelectionData) {}
