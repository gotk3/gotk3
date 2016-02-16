package gtk

import "github.com/gotk3/gotk3/glib"

type FileFilter interface {
	glib.Object

	AddPattern(string)
	AddPixbufFormats()
	SetName(string)
} // end of FileFilter

func AssertFileFilter(_ FileFilter) {}
