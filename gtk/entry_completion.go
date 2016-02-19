package gtk

import "github.com/gotk3/gotk3/glib"

type EntryCompletion interface {
	glib.Object
} // end of EntryCompletion

func AssertEntryCompletion(_ EntryCompletion) {}
