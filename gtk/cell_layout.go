package gtk

import "github.com/gotk3/gotk3/glib"

type CellLayout interface {
	glib.Object

	AddAttribute(CellRenderer, string, int)
	PackStart(CellRenderer, bool)
} // end of CellLayout

func AssertCellLayout(_ CellLayout) {}
