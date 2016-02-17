package gtk

import "github.com/gotk3/gotk3/glib"

type Builder interface {
	glib.Object

	AddFromFile(string) error
	AddFromResource(string) error
	AddFromString(string) error
	ConnectSignals(map[string]interface{})
	GetObject(string) (glib.Object, error)
} // end of Builder

func AssertBuilder(_ Builder) {}
