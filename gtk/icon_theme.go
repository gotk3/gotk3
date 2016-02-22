package gtk

import "github.com/gotk3/gotk3/gdk"

type IconTheme interface {
	LoadIcon(string, int, IconLookupFlags) (gdk.Pixbuf, error)
} // end of IconTheme

func AssertIconTheme(_ IconTheme) {}
