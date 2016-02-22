package gtk

import "github.com/gotk3/gotk3/glib"

type TreeSelection interface {
	glib.Object

	CountSelectedRows() int
	GetMode() SelectionMode
	GetSelected() (TreeModel, TreeIter, bool)
	GetSelectedRows(TreeModel) glib.List
	SelectIter(TreeIter)
	SelectPath(TreePath)
	SetMode(SelectionMode)
	UnselectPath(TreePath)
} // end of TreeSelection

func AssertTreeSelection(_ TreeSelection) {}
