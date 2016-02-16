package gtk

import "github.com/gotk3/gotk3/glib"

type Orientable interface {
	glib.Object

	GetOrientation() Orientation
	SetOrientation(Orientation)
} // end of Orientable

func AssertOrientable(_ Orientable) {}
