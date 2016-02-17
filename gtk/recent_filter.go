package gtk

import "github.com/gotk3/gotk3/glib"

type RecentFilter interface {
	glib.InitiallyUnowned
} // end of RecentFilter

func AssertRecentFilter(_ RecentFilter) {}
