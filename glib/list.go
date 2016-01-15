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
	// If set, dataWrap is called every time NthDataWrapped()
	// or DataWrapped() is called to wrap raw underlying
	// value into appropriate type.
	dataWrap func(unsafe.Pointer) interface{}
}

func WrapList(obj uintptr) *List {
	return wrapList((*C.struct__GList)(unsafe.Pointer(obj)))
}

func wrapList(obj *C.struct__GList) *List {
	if obj == nil {
		return nil
	}
	return &List{list: obj}
}

func (v *List) wrapNewHead(obj *C.struct__GList) *List {
	if obj == nil {
		return nil
	}
	return &List{
		list:     obj,
		dataWrap: v.dataWrap,
	}
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

// DataWapper sets wrap functions, which is called during NthDataWrapped()
// and DataWrapped(). It's used to cast raw C data into appropriate
// Go structures and types every time that data is retreived.
func (v *List) DataWrapper(fn func(unsafe.Pointer) interface{}) {
	v.dataWrap = fn
}

// Append is a wrapper around g_list_append().
func (v *List) Append(data uintptr) *List {
	glist := C.g_list_append(v.native(), C.gpointer(data))
	return v.wrapNewHead(glist)
}

// Prepend is a wrapper around g_list_prepend().
func (v *List) Prepend(data uintptr) *List {
	glist := C.g_list_prepend(v.native(), C.gpointer(data))
	return v.wrapNewHead(glist)
}

// Insert is a wrapper around g_list_insert().
func (v *List) Insert(data uintptr, position int) *List {
	glist := C.g_list_insert(v.native(), C.gpointer(data), C.gint(position))
	return v.wrapNewHead(glist)
}

// Length is a wrapper around g_list_length().
func (v *List) Length() uint {
	return uint(C.g_list_length(v.native()))
}

// NthData is a wrapper around g_list_nth_data().
func (v *List) NthData(n uint) unsafe.Pointer {
	return unsafe.Pointer(C.g_list_nth_data(v.native(), C.guint(n)))
}

// NthDataWrapped acts the same as NthData(), but passes
// retrieved value before returning through wrap function, set by DataWrapper().
// If no wrap function is set, it returns raw unsafe.Pointer.
func (v *List) NthDataWrapped(n uint) interface{} {
	ptr := v.NthData(n)
	if v.dataWrap != nil {
		return v.dataWrap(ptr)
	}
	return ptr
}

// Free is a wrapper around g_list_free().
func (v *List) Free() {
	C.g_list_free(v.native())
}

// Next is a wrapper around the next struct field
func (v *List) Next() *List {
	return v.wrapNewHead(v.native().next)
}

// Previous is a wrapper around the prev struct field
func (v *List) Previous() *List {
	return v.wrapNewHead(v.native().prev)
}

// Data is a wrapper around the data struct field
func (v *List) Data() unsafe.Pointer {
	return unsafe.Pointer(v.native().data)
}

// DataWrapped acts the same as Data(), but passes
// retrieved value before returning through wrap function, set by DataWrapper().
// If no wrap function is set, it returns raw unsafe.Pointer.
func (v *List) DataWrapped() interface{} {
	ptr := v.Data()
	if v.dataWrap != nil {
		return v.dataWrap(ptr)
	}
	return ptr
}
