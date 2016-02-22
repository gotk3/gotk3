package gdk

import "github.com/gotk3/gotk3/glib"

type Cursor interface {
	glib.Object
} // end of Cursor

func AssertCursor(_ Cursor) {}
