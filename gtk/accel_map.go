package gtk

import "github.com/gotk3/gotk3/glib"

type AccelMap interface {
	glib.Object
} // end of AccelMap

func AssertAccelMap(_ AccelMap) {}
