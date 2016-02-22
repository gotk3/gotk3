package gdk

type Rectangle interface {
	GetHeight() int
	GetWidth() int
	GetX() int
	GetY() int
} // end of Rectangle

func AssertRectangle(_ Rectangle) {}
