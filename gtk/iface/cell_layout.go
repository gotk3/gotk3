package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type CellLayout interface {
	glib_iface.Object

	AddAttribute(CellRenderer, string, int)
	PackStart(CellRenderer, bool)
} // end of CellLayout

func AssertCellLayout(_ CellLayout) {}
