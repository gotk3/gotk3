package gtk

import "github.com/gotk3/gotk3/glib"

type CellRenderer interface {
	glib.InitiallyUnowned
} // end of CellRenderer

func AssertCellRenderer(_ CellRenderer) {}
