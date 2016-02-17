package pango

type AttrShape interface {
	Attribute
} // end of AttrShape

func AssertAttrShape(_ AttrShape) {}
