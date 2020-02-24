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
	toGListModel() *C.GListModel
}

// ListModel is a representation of GIO's GListModel.
type ListModel struct {
	*Object
}

func (v *ListModel) toGListModel() *C.GListModel {
	if v == nil {
		return nil
	}
	return v.native()
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

func marshalListModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapListModel(wrapObject(unsafe.Pointer(c))), nil
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

/*
 * GListStore
 */

// ListStore is a representation of GListStore
type ListStore struct {
	ListModel
}

func (v *ListStore) native() *C.GListStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGListStore(unsafe.Pointer(v.GObject))
}

func (v *ListStore) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalListStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapListStore(wrapObject(unsafe.Pointer(c))), nil
}

func wrapListStore(obj *Object) *ListStore {
	return &ListStore{ListModel{obj}}
}

// ListStoreNew is a wrapper around g_list_store_new().
func ListStoreNew(itemType Type) *ListStore {
	c := C.g_list_store_new(C.GType(itemType))
	if c == nil {
		return nil
	}
	return wrapListStore(wrapObject(unsafe.Pointer(c)))
}

// Insert is a wrapper around g_list_store_insert().
func (v *ListStore) Insert(position uint, item uintptr) {
	C.g_list_store_insert(v.native(), C.guint(position), C.gpointer(item))
}

// InsertSorted is a wrapper around g_list_store_insert_sorted().
// func (v *ListStore) InsertSorted(item uintptr, compareFunc CompareDataFunc, userData ...interface{}) {
// }

// Append is a wrapper around g_list_store_append().
func (v *ListStore) Append(item uintptr) {
	C.g_list_store_append(v.native(), C.gpointer(item))
}

// Remove is a wrapper around g_list_store_remove().
func (v *ListStore) Remove(position uint) {
	C.g_list_store_remove(v.native(), C.guint(position))
}

// Splice is a wrapper around g_list_store_splice().
func (v *ListStore) Splice(position uint, removalLength uint, additions []interface{}) {

	additionsLength := len(additions)
	gAdditions := make([]*C.GObject, additionsLength)
	for i, add := range additions {
		gAdditions[i] = ToGObject(unsafe.Pointer(&add))
	}
	gAdditions = append(gAdditions, nil)

	C.g_list_store_splice(v.native(), C.guint(position), C.guint(removalLength), &additions[0], C.guint(additionsLength))
}

// Sort is a wrapper around g_list_store_sort().
// func (v *ListStore) Sort(compareFunc CompareDataFunc, userData ...interface{}) {
// }
