package gtk

import "github.com/gotk3/gotk3/glib"

type TreeStore interface {
	glib.Object

	Append(TreeIter) TreeIter
	Clear()
	Insert(TreeIter, int) TreeIter
	Remove(TreeIter) bool
	SetValue(TreeIter, int, interface{}) error
} // end of TreeStore

func AssertTreeStore(_ TreeStore) {}
