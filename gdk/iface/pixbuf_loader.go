package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type PixbufLoader interface {
    glib_iface.Object

    Close() error
    GetPixbuf() (Pixbuf, error)
    SetSize(int, int)
    Write([]byte) (int, error)
} // end of PixbufLoader

func AssertPixbufLoader(_ PixbufLoader) {}
