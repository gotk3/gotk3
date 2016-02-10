package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RecentChooser interface {
    glib_iface.Object

    AddFilter(RecentFilter)
    GetCurrentUri() string
    RemoveFilter(RecentFilter)
} // end of RecentChooser

func AssertRecentChooser(_ RecentChooser) {}
