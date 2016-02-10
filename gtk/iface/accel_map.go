package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type AccelMap interface {
    glib_iface.Object
} // end of AccelMap

func AssertAccelMap(_ AccelMap) {}
