package iface

type Box interface {
	Container

	GetHomogeneous() bool
	GetSpacing() int
	PackEnd(Widget, bool, bool, uint)
	PackStart(Widget, bool, bool, uint)
	QueryChildPacking(Widget) (bool, bool, uint, PackType)
	ReorderChild(Widget, int)
	SetChildPacking(Widget, bool, bool, uint, PackType)
	SetHomogeneous(bool)
	SetSpacing(int)
} // end of Box

func AssertBox(_ Box) {}
