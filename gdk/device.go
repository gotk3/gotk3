package gdk

import "github.com/gotk3/gotk3/glib"

type Device interface {
	glib.Object

	Grab(Window, GrabOwnership, bool, EventMask, Cursor, uint32) GrabStatus
	Ungrab(uint32)
} // end of Device

func AssertDevice(_ Device) {}
