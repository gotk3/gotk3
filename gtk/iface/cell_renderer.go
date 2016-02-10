package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type CellRenderer interface {
    glib_iface.InitiallyUnowned
} // end of CellRenderer

func AssertCellRenderer(_ CellRenderer) {}
