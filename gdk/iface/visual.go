package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Visual interface {
    glib_iface.Object
} // end of Visual

func AssertVisual(_ Visual) {}
