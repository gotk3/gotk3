package gdk

import "github.com/gotk3/gotk3/glib"

type Visual interface {
	glib.Object
} // end of Visual

func AssertVisual(_ Visual) {}
