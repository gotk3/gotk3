package gtk

import "github.com/gotk3/gotk3/glib"

type CssProvider interface {
	glib.Object

	LoadFromData(string) error
	LoadFromPath(string) error
	ToString() (string, error)
} // end of CssProvider

func AssertCssProvider(_ CssProvider) {}
