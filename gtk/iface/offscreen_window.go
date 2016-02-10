package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"
import cairo_iface "github.com/gotk3/gotk3/cairo/iface"

type OffscreenWindow interface {
    Window

    GetPixbuf() (gdk_iface.Pixbuf, error)
    GetSurface() (cairo_iface.Surface, error)
} // end of OffscreenWindow

func AssertOffscreenWindow(_ OffscreenWindow) {}
