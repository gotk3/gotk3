package gdk

type EventButton interface {
	Event

	Button() uint
	ButtonVal() uint
	MotionVal() (float64, float64)
	MotionValRoot() (float64, float64)
	State() uint
	Time() uint32
	Type() EventType
	X() float64
	XRoot() float64
	Y() float64
	YRoot() float64
} // end of EventButton

func AssertEventButton(_ EventButton) {}
