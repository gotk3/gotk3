package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Window interface {
	glib_iface.Object

	GetDesktop() uint32
	MoveToCurrentDesktop()
	MoveToDesktop(uint32)
} // end of Window

func AssertWindow(_ Window) {}
