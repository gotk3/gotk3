package gtk

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
)

type OffscreenWindow interface {
	Window

	GetPixbuf() (gdk.Pixbuf, error)
	GetSurface() (cairo.Surface, error)
} // end of OffscreenWindow

func AssertOffscreenWindow(_ OffscreenWindow) {}
