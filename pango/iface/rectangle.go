package iface


type Rectangle interface {
    ExtentsToPixels(Rectangle)
} // end of Rectangle

func AssertRectangle(_ Rectangle) {}
