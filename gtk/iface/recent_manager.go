package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RecentManager interface {
    glib_iface.Object

    AddItem(string) bool
} // end of RecentManager

func AssertRecentManager(_ RecentManager) {}
