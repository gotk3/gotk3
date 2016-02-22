package gtk

type Bin interface {
	Container

	GetChild() (Widget, error)
} // end of Bin

func AssertBin(_ Bin) {}
