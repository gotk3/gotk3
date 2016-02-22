package gtk

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

type ColorChooser interface {
	glib.Object

	AddPalette(Orientation, int, []gdk.RGBA)
	GetRGBA() gdk.RGBA
	GetUseAlpha() bool
	SetRGBA(gdk.RGBA)
	SetUseAlpha(bool)
} // end of ColorChooser

func AssertColorChooser(_ ColorChooser) {}
