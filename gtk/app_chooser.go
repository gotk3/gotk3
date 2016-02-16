package gtk

import "github.com/gotk3/gotk3/glib"

type AppChooser interface {
	glib.Object

	GetContentType() string
	Refresh()
} // end of AppChooser

func AssertAppChooser(_ AppChooser) {}
