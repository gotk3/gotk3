package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Orientable interface {
    glib_iface.Object

    GetOrientation() Orientation
    SetOrientation(Orientation)
} // end of Orientable

func AssertOrientable(_ Orientable) {}
