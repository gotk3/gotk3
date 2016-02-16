package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type TreeSelection interface {
	glib_iface.Object

	CountSelectedRows() int
	GetMode() SelectionMode
	GetSelected() (TreeModel, TreeIter, bool)
	GetSelectedRows(TreeModel) glib_iface.List
	SelectIter(TreeIter)
	SelectPath(TreePath)
	SetMode(SelectionMode)
	UnselectPath(TreePath)
} // end of TreeSelection

func AssertTreeSelection(_ TreeSelection) {}
