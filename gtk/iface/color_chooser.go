package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"
import glib_iface "github.com/gotk3/gotk3/glib/iface"

type ColorChooser interface {
    glib_iface.Object

    AddPalette(Orientation, int, []gdk_iface.RGBA)
    GetRGBA() gdk_iface.RGBA
    GetUseAlpha() bool
    SetRGBA(gdk_iface.RGBA)
    SetUseAlpha(bool)
} // end of ColorChooser

func AssertColorChooser(_ ColorChooser) {}
