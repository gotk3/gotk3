package gtk

import "github.com/gotk3/gotk3/gdk"

type Image interface {
	Widget

	Clear()
	GetIconName() (string, IconSize)
	GetPixbuf() gdk.Pixbuf
	GetPixelSize() int
	GetStorageType() ImageType
	SetFromFile(string)
	SetFromIconName(string, IconSize)
	SetFromPixbuf(gdk.Pixbuf)
	SetFromResource(string)
	SetPixelSize(int)
} // end of Image

func AssertImage(_ Image) {}
