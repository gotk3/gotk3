package gtk

type DrawingArea interface {
	Widget
} // end of DrawingArea

func AssertDrawingArea(_ DrawingArea) {}
