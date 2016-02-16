package gtk

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

type TextTag interface {
	glib.Object

	Event(glib.Object, gdk.Event, TextIter) bool
	GetPriority() int
	SetPriority(int)
} // end of TextTag

func AssertTextTag(_ TextTag) {}
