package gtk

import "github.com/gotk3/gotk3/glib"

type RecentChooser interface {
	glib.Object

	AddFilter(RecentFilter)
	GetCurrentUri() string
	RemoveFilter(RecentFilter)
} // end of RecentChooser

func AssertRecentChooser(_ RecentChooser) {}
