package gtk

type Layout interface {
	Container

	GetSize() (uint, uint)
	Move(Widget, int, int)
	Put(Widget, int, int)
	SetSize(uint, uint)
} // end of Layout

func AssertLayout(_ Layout) {}
