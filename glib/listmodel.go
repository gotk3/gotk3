package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
 * GListModel
 */

// IListModel is an interface representation of ListModel,
// used to avoid duplication when embedding the type in a wrapper of another GObject-based type.
type IListModel interface {
	GetItemType() Type
	GetNItems() uint
	GetItem(position uint) uintptr
	GetObject(position uint) *Object
	ItemsChanged(position, removed, added uint)
}

// ListModel is a representation of GIO's GListModel.
type ListModel struct {
	*Object
}

// native returns a pointer to the underlying GListModel.
func (v *ListModel) native() *C.GListModel {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGListModel(unsafe.Pointer(v.GObject))
}

// Native returns a pointer to the underlying GListModel.
func (v *ListModel) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func wrapListModel(obj *Object) *ListModel {
	return &ListModel{obj}
}

// GetItemType is a wrapper around g_list_model_get_item_type().
func (v *ListModel) GetItemType() Type {
	return Type(C.g_list_model_get_item_type(v.native()))
}

// GetNItems is a wrapper around g_list_model_get_n_items().
func (v *ListModel) GetNItems() uint {
	return int(C.g_list_model_get_n_items(v.native()))
}

// GetItem is a wrapper around g_list_model_get_item().
func (v *ListModel) GetItem(position uint) uintptr {
	c := C.g_list_model_get_item(v.native(), C.guint(position))
	return uintptr(unsafe.Pointer(c))
}

// GetObject is a wrapper around g_list_model_get_object().
func (v *ListModel) GetObject(position uint) *Object {
	c := C.g_list_model_get_object(v.native(), C.guint(position))
	return wrapObject(unsafe.Pointer(c))
}

// ItemsChanged is a wrapper around g_list_model_items_changed().
func (v *ListModel) ItemsChanged(position, removed, added int) {
	C.g_list_model_items_changed(v.native(), uint(position), uint(removed), uint(added))
}
