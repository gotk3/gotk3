package gtk

import "github.com/gotk3/gotk3/glib"

type ListStore interface {
	glib.Object

	Append() TreeIter
	Clear()
	InsertAfter(TreeIter) TreeIter
	InsertBefore(TreeIter) TreeIter
	IterIsValid(TreeIter) bool
	MoveAfter(TreeIter, TreeIter)
	MoveBefore(TreeIter, TreeIter)
	Prepend() TreeIter
	Remove(TreeIter) bool
	Set2(TreeIter, []int, []interface{}) error
	SetCols(TreeIter, Cols) error
	SetSortColumnId(int, SortType)
	SetValue(TreeIter, int, interface{}) error
	Swap(TreeIter, TreeIter)
} // end of ListStore

func AssertListStore(_ ListStore) {}
