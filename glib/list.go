package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
 * Linked Lists
 */

// List is a representation of Glib's GList.
type List struct {
	list *C.struct__GList
}

func WrapList(obj uintptr) *List {
	return wrapList((*C.struct__GList)(unsafe.Pointer(obj)))
}

func wrapList(obj *C.struct__GList) *List {
	return &List{obj}
}

func (v *List) Native() uintptr {
	return uintptr(unsafe.Pointer(v.list))
}

func (v *List) native() *C.struct__GList {
	if v == nil || v.list == nil {
		return nil
	}
	return v.list
}

// Append is a wrapper around g_list_append().
func (v *List) Append(data uintptr) *List {
	glist := C.g_list_append(v.native(), C.gpointer(data))
	return wrapList(glist)
}

// Prepend is a wrapper around g_list_prepend().
func (v *List) Prepend(data uintptr) *List {
	glist := C.g_list_prepend(v.native(), C.gpointer(data))
	return wrapList(glist)
}

// Insert is a wrapper around g_list_insert().
func (v *List) Insert(data uintptr, position int) *List {
	glist := C.g_list_insert(v.native(), C.gpointer(data), C.gint(position))
	return wrapList(glist)
}

// Length is a wrapper around g_list_length().
func (v *List) Length() uint {
	return uint(C.g_list_length(v.native()))
}

// NthData is a wrapper around g_list_nth_data().
func (v *List) NthData(n uint) interface{} {
	return C.g_list_nth_data(v.native(), C.guint(n))
}

// Free is a wrapper around g_list_free().
func (v *List) Free() {
	C.g_list_free(v.native())
}

// Next is a wrapper around the next struct field
func (v *List) Next() *List {
	return wrapList(v.native().next)
}

// Previous is a wrapper around the prev struct field
func (v *List) Previous() *List {
	return wrapList(v.native().prev)
}

// Data is a wrapper around the data struct field
func (v *List) Data() uintptr {
	return uintptr(v.native().data)
}
