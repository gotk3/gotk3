package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type CssProvider interface {
    glib_iface.Object

    LoadFromData(string) error
    LoadFromPath(string) error
    ToString() (string, error)
} // end of CssProvider

func AssertCssProvider(_ CssProvider) {}
