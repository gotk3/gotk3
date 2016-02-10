package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type AppChooser interface {
    glib_iface.Object

    GetContentType() string
    Refresh()
} // end of AppChooser

func AssertAppChooser(_ AppChooser) {}
