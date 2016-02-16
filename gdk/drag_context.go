package gdk

import "github.com/gotk3/gotk3/glib"

type DragContext interface {
	glib.Object

	ListTargets() glib.List
} // end of DragContext

func AssertDragContext(_ DragContext) {}
