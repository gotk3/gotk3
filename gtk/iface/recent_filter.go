package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RecentFilter interface {
    glib_iface.InitiallyUnowned
} // end of RecentFilter

func AssertRecentFilter(_ RecentFilter) {}
