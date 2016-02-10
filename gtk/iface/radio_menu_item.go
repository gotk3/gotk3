package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RadioMenuItem interface {
    CheckMenuItem

    GetGroup() (glib_iface.SList, error)
    SetGroup(glib_iface.SList)
} // end of RadioMenuItem

func AssertRadioMenuItem(_ RadioMenuItem) {}
