package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"unsafe"
)

// IAsyncResult is an interface representation of AsyncResult,
// used to avoid duplication when embedding the type in a wrapper of another GObject-based type.
type IAsyncResult interface {
	GetUserData() uintptr
	GetSourceObject() *Object
	IsTagged(sourceTag uintptr) bool
	LegacyPropagateError() error
}

// AsyncReadyCallback is a representation of GAsyncReadyCallback
type AsyncReadyCallback func(object *Object, res *AsyncResult)

// AsyncResult is a representation of GIO's GAsyncResult.
type AsyncResult struct {
	*Object
}

// native() returns a pointer to the underlying GAsyncResult.
func (v *AsyncResult) native() *C.GAsyncResult {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGAsyncResult(unsafe.Pointer(v.GObject))
}

func wrapAsyncResult(obj *Object) *AsyncResult {
	return &AsyncResult{obj}
}

// GetUserData is a wrapper around g_async_result_get_user_data()
func (v *AsyncResult) GetUserData() uintptr {
	c := C.g_async_result_get_user_data(v.native())
	return uintptr(unsafe.Pointer(c))
}

// GetSourceObject is a wrapper around g_async_result_get_source_object
func (v *AsyncResult) GetSourceObject() *Object {
	obj := C.g_async_result_get_source_object(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

// IsTagged is a wrapper around g_async_result_is_tagged
func (v *AsyncResult) IsTagged(sourceTag uintptr) bool {
	c := C.g_async_result_is_tagged(v.native(), C.gpointer(sourceTag))
	return gobool(c)
}

// LegacyPropagateError is a wrapper around g_async_result_legacy_propagate_error
func (v *AsyncResult) LegacyPropagateError() error {
	var err *C.GError
	c := C.g_async_result_legacy_propagate_error(v.native(), &err)
	isSimpleAsyncResult := gobool(c)
	if isSimpleAsyncResult {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}
