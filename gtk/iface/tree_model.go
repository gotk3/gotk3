package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type TreeModel interface {
    glib_iface.Object

    GetColumnType(int) glib_iface.Type
    GetFlags() TreeModelFlags
    GetIter(TreePath) (TreeIter, error)
    GetIterFirst() (TreeIter, bool)
    GetIterFromString(string) (TreeIter, error)
    GetNColumns() int
    GetPath(TreeIter) (TreePath, error)
    GetValue(TreeIter, int) (glib_iface.Value, error)
    IterChildren(TreeIter, TreeIter) bool
    IterNChildren(TreeIter) int
    IterNext(TreeIter) bool
    IterPrevious(TreeIter) bool
} // end of TreeModel

func AssertTreeModel(_ TreeModel) {}
