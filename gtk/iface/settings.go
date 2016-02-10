package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Settings interface {
    glib_iface.Object
} // end of Settings

func AssertSettings(_ Settings) {}
