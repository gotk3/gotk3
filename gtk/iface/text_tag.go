package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"
import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type TextTag interface {
	glib_iface.Object

	Event(glib_iface.Object, gdk_iface.Event, TextIter) bool
	GetPriority() int
	SetPriority(int)
} // end of TextTag

func AssertTextTag(_ TextTag) {}
