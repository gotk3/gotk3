package gtk

type EventBox interface {
	Bin

	GetAboveChild() bool
	GetVisibleWindow() bool
	SetAboveChild(bool)
	SetVisibleWindow(bool)
} // end of EventBox

func AssertEventBox(_ EventBox) {}
