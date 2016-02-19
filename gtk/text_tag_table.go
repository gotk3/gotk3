package gtk

import "github.com/gotk3/gotk3/glib"

type TextTagTable interface {
	glib.Object

	Add(TextTag)
	Lookup(string) (TextTag, error)
	Remove(TextTag)
} // end of TextTagTable

func AssertTextTagTable(_ TextTagTable) {}
