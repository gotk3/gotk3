// package resource wraps operations over GResource
package gio

// #cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <stdlib.h>
// #include "gresource.go.h"
import "C"
import (
	"errors"
	"unsafe"
)

// ResourceLookupFlags is a representation of GTK's GResourceLookupFlags
type ResourceLookupFlags int

func (f ResourceLookupFlags) native() C.GResourceLookupFlags {
	return (C.GResourceLookupFlags)(f)
}

const (
	G_RESOURCE_LOOKUP_FLAGS_NONE ResourceLookupFlags = C.G_RESOURCE_LOOKUP_FLAGS_NONE
)

// GResource wraps native GResource object
//
// See: https://developer.gnome.org/gio/stable/GResource.html
type GResource *C.GResource

// LoadGResource is a wrapper around g_resource_load()
//
// See: https://developer.gnome.org/gio/stable/GResource.html#g-resource-load
func LoadGResource(path string) (GResource, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	var gerr *C.GError

	resPtr := C.g_resource_load((*C.gchar)(unsafe.Pointer(cpath)), &gerr)
	if gerr != nil {
		defer C.g_error_free(gerr)
		return nil, errors.New(goString(gerr.message))
	}

	res := wrapGResource(resPtr)
	return res, nil
}

// NewGResourceFromData is a wrapper around g_resource_new_from_data()
//
// See: https://developer.gnome.org/gio/stable/GResource.html#g-resource-new-from-data
func NewGResourceFromData(data []byte) (GResource, error) {
	arrayPtr := (*C.GBytes)(unsafe.Pointer(&data[0]))
	var gerr *C.GError
	resPtr := C.g_resource_new_from_data(arrayPtr, &gerr)
	if gerr != nil {
		defer C.g_error_free(gerr)
		return nil, errors.New(goString(gerr.message))
	}

	res := wrapGResource(resPtr)
	return res, nil
}

// Register wraps g_resources_register()
//
// See: https://developer.gnome.org/gio/stable/GResource.html#g-resources-register
func RegisterGResource(res GResource) {
	C.g_resources_register(res)
}

// Unregister wraps g_resources_unregister()
//
// See: https://developer.gnome.org/gio/stable/GResource.html#g-resources-unregister
func UnregisterGResource(res GResource) {
	C.g_resources_unregister(res)
}

// GResourceEnumerateChildren wraps g_resources_enumerate_children()
//
// See: https://developer.gnome.org/gio/stable/GResource.html#g-resources-enumerate-children
func GResourceEnumerateChildren(path string, flags ResourceLookupFlags) ([]string, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	var gerr *C.GError
	arrChildren := C.g_resources_enumerate_children(cpath, flags.native(), &gerr)
	if gerr != nil {
		defer C.g_error_free(gerr)
		return nil, errors.New(goString(gerr.message))
	}

	if arrChildren == nil {
		return nil, errors.New("unexpected nil pointer from g_resources_enumerate_children")
	}

	arr := toGoStringArray(arrChildren)
	return arr, nil
}

func wrapGResource(resPtr *C.GResource) GResource {
	res := GResource(resPtr)
	return res
}
