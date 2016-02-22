package pango

type AttrColor interface {
	Attribute
	Color
} // end of AttrColor

func AssertAttrColor(_ AttrColor) {}
