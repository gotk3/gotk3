package gdk

type EventScroll interface {
	Event

	DeltaX() float64
	DeltaY() float64
	Type() EventType
	X() float64
	Y() float64
} // end of EventScroll

func AssertEventScroll(_ EventScroll) {}
