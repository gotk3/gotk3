package iface

type Paned interface {
	Bin

	Add1(Widget)
	Add2(Widget)
	GetChild1() (Widget, error)
	GetChild2() (Widget, error)
	GetHandleWindow() (Window, error)
	GetPosition() int
	Pack1(Widget, bool, bool)
	Pack2(Widget, bool, bool)
	SetPosition(int)
} // end of Paned

func AssertPaned(_ Paned) {}
