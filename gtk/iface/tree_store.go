package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type TreeStore interface {
	glib_iface.Object

	Append(TreeIter) TreeIter
	Clear()
	Insert(TreeIter, int) TreeIter
	Remove(TreeIter) bool
	SetValue(TreeIter, int, interface{}) error
} // end of TreeStore

func AssertTreeStore(_ TreeStore) {}
