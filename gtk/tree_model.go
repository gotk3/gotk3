package gtk

import "github.com/gotk3/gotk3/glib"

type TreeModel interface {
	glib.Object

	GetColumnType(int) glib.Type
	GetFlags() TreeModelFlags
	GetIter(TreePath) (TreeIter, error)
	GetIterFirst() (TreeIter, bool)
	GetIterFromString(string) (TreeIter, error)
	GetNColumns() int
	GetPath(TreeIter) (TreePath, error)
	GetValue(TreeIter, int) (glib.Value, error)
	IterChildren(TreeIter, TreeIter) bool
	IterNChildren(TreeIter) int
	IterNext(TreeIter) bool
	IterPrevious(TreeIter) bool
} // end of TreeModel

func AssertTreeModel(_ TreeModel) {}
