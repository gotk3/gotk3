package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type ListStore interface {
	glib_iface.Object

	Append() TreeIter
	Clear()
	InsertAfter(TreeIter) TreeIter
	InsertBefore(TreeIter) TreeIter
	IterIsValid(TreeIter) bool
	MoveAfter(TreeIter, TreeIter)
	MoveBefore(TreeIter, TreeIter)
	Prepend() TreeIter
	Remove(TreeIter) bool
	SetCols(TreeIter, Cols) error
	SetSortColumnId(int, SortType)
	SetValue(TreeIter, int, interface{}) error
	Swap(TreeIter, TreeIter)
} // end of ListStore

func AssertListStore(_ ListStore) {}
