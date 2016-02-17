package gdk

import "github.com/gotk3/gotk3/glib"

type PixbufLoader interface {
	glib.Object

	Close() error
	GetPixbuf() (Pixbuf, error)
	SetSize(int, int)
	Write([]byte) (int, error)
} // end of PixbufLoader

func AssertPixbufLoader(_ PixbufLoader) {}
