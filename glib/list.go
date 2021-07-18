package glib

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"runtime"
	"unsafe"
)

/*
 * Linked Lists
 */

// List is a representation of Glib's GList.
type List struct {
	list *C.struct__GList
	// If set, dataWrap is called every time NthData()
	// or Data() is called to wrap raw underlying
	// value into appropriate type.
	dataWrap func(unsafe.Pointer) interface{}
	head     *List // keep the head alive, prevent GC
}

func WrapList(obj uintptr) *List {
	return wrapList((*C.struct__GList)(unsafe.Pointer(obj)))
}

func wrapList(obj *C.struct__GList) *List {
	if obj == nil {
		return nil
	}
	l := &List{list: obj}
	l.head = l
	return l
}

func (v *List) wrapNewHead(obj *C.struct__GList) *List {
	if obj == nil {
		return nil
	}
	return &List{
		list:     obj,
		dataWrap: v.dataWrap,
		head:     v,
	}
}

// AttachFinalizer attaches a finalizer into the List to be freed when it is no
// longer used by Go. New List heads created from this list will have the
// finalizer attached as well. If the freeData callback is not nil, then the
// data inside will be freed as well.
//
// It is the caller's responsibility to call this method on the head of the list
// and not on any of the children heads. Doing so is undefined behavior and may
// cause double frees.
func (v *List) AttachFinalizer(freeData func(uintptr)) {
	if v == nil {
		return
	}

	runtime.SetFinalizer(v, func(v *List) {
		if freeData != nil {
			for list := v.list; list != nil; list = list.next {
				freeData(uintptr(list.data))
			}
		}

		C.g_list_free(v.list)
	})
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

// DataWapper sets wrap functions, which is called during NthData()
// and Data(). It's used to cast raw C data into appropriate
// Go structures and types every time that data is retreived.
func (v *List) DataWrapper(fn func(unsafe.Pointer) interface{}) {
	if v == nil {
		return
	}
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

// nthDataRaw is a wrapper around g_list_nth_data().
func (v *List) nthDataRaw(n uint) unsafe.Pointer {
	return unsafe.Pointer(C.g_list_nth_data(v.native(), C.guint(n)))
}

// Nth() is a wrapper around g_list_nth().
func (v *List) Nth(n uint) *List {
	list := wrapList(C.g_list_nth(v.native(), C.guint(n)))
	list.DataWrapper(v.dataWrap)
	return list
}

// NthData acts the same as g_list_nth_data(), but passes
// retrieved value before returning through wrap function, set by DataWrapper().
// If no wrap function is set, it returns raw unsafe.Pointer.
func (v *List) NthData(n uint) interface{} {
	ptr := v.nthDataRaw(n)
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

// First is a wrapper around g_list_first().
func (v *List) First() *List {
	return v.wrapNewHead(C.g_list_first(v.native()))
}

// Last is a wrapper around g_list_last().
func (v *List) Last() *List {
	return v.wrapNewHead(C.g_list_last(v.native()))
}

// Reverse is a wrapper around g_list_reverse().
func (v *List) Reverse() *List {
	return v.wrapNewHead(C.g_list_reverse(v.native()))
}

// dataRaw is a wrapper around the data struct field
func (v *List) dataRaw() unsafe.Pointer {
	return unsafe.Pointer(v.native().data)
}

// Data acts the same as data struct field, but passes
// retrieved value before returning through wrap function, set by DataWrapper().
// If no wrap function is set, it returns raw unsafe.Pointer.
func (v *List) Data() interface{} {
	ptr := v.dataRaw()
	if v.dataWrap != nil {
		return v.dataWrap(ptr)
	}
	return ptr
}

// Foreach acts the same as g_list_foreach().
// No user_data argument is implemented because of Go clojure capabilities.
func (v *List) Foreach(fn func(item interface{})) {
	for l := v; l != nil; l = l.Next() {
		fn(l.Data())
	}
}

// FreeFull acts the same as g_list_free_full().
// Calling list.FreeFull(fn) is equivalent to calling list.Foreach(fn) and
// list.Free() sequentially.
func (v *List) FreeFull(fn func(item interface{})) {
	v.Foreach(fn)
	v.Free()
}

// CompareDataFunc is a representation of GCompareDataFunc
type CompareDataFunc func(a, b uintptr) int
