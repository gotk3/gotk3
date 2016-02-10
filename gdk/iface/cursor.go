package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Cursor interface {
    glib_iface.Object
} // end of Cursor

func AssertCursor(_ Cursor) {}
