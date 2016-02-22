package gtk

import "github.com/gotk3/gotk3/glib"

type Settings interface {
	glib.Object
} // end of Settings

func AssertSettings(_ Settings) {}
