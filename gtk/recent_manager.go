package gtk

import "github.com/gotk3/gotk3/glib"

type RecentManager interface {
	glib.Object

	AddItem(string) bool
} // end of RecentManager

func AssertRecentManager(_ RecentManager) {}
