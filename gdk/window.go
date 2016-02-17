package gdk

import "github.com/gotk3/gotk3/glib"

type Window interface {
	glib.Object

	GetDesktop() uint32
	MoveToCurrentDesktop()
	MoveToDesktop(uint32)
} // end of Window

func AssertWindow(_ Window) {}
