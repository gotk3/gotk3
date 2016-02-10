package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type IconTheme interface {
    LoadIcon(string, int, IconLookupFlags) (gdk_iface.Pixbuf, error)
} // end of IconTheme

func AssertIconTheme(_ IconTheme) {}
