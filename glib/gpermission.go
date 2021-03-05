package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "gpermission.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/internal/callback"
)

// Permission is a representation of GIO's GPermission.
type Permission struct {
	*Object
}

func (v *Permission) native() *C.GPermission {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGPermission(unsafe.Pointer(v.GObject))
}

// Native returns a uintptr to the underlying C.GPermission.
func (v *Permission) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalPermission(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPermission(wrapObject(unsafe.Pointer(c))), nil
}

func wrapPermission(obj *Object) *Permission {
	return &Permission{obj}
}

// WrapPermission wraps given unsafe pointer into Permission.
func WrapPermission(ptr unsafe.Pointer) *Permission {
	return wrapPermission(wrapObject(ptr))
}

// GetAllowed is a wrapper around g_permission_get_allowed().
func (v *Permission) GetAllowed() bool {
	c := C.g_permission_get_allowed(v.native())
	return gobool(c)
}

// GetCanAcquire is a wrapper around g_permission_get_can_acquire().
func (v *Permission) GetCanAcquire() bool {
	c := C.g_permission_get_can_acquire(v.native())
	return gobool(c)
}

// GetCanRelease is a wrapper around g_permission_get_can_release().
func (v *Permission) GetCanRelease() bool {
	c := C.g_permission_get_can_release(v.native())
	return gobool(c)
}

// Acquire is a wrapper around g_permission_acquire().
func (v *Permission) Acquire(cancellable *Cancellable) error {
	var err *C.GError
	c := C.g_permission_acquire(v.native(), cancellable.native(), &err)
	acquired := gobool(c)
	if !acquired {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// AcquireAsync is a wrapper around g_permission_acquire_async().
func (v *Permission) AcquireAsync(cancellable *Cancellable, fn AsyncReadyCallback) {
	C._g_permission_acquire_async(v.native(), cancellable.native(), C.gpointer(callback.Assign(fn)))
}

// AcquireFinish is a wrapper around g_permission_acquire_finish().
func (v *Permission) AcquireFinish(result *AsyncResult) error {
	var err *C.GError
	c := C.g_permission_acquire_finish(v.native(), result.native(), &err)
	acquired := gobool(c)
	if !acquired {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// Release is a wrapper around g_permission_release().
func (v *Permission) Release(cancellable *Cancellable) error {
	var err *C.GError
	c := C.g_permission_release(v.native(), cancellable.native(), &err)
	released := gobool(c)
	if !released {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// ReleaseAsync is a wrapper around g_permission_release_async().
func (v *Permission) ReleaseAsync(cancellable *Cancellable, fn AsyncReadyCallback) {
	C._g_permission_release_async(v.native(), cancellable.native(), C.gpointer(callback.Assign(fn)))
}

// ReleaseFinish is a wrapper around g_permission_release_finish().
func (v *Permission) ReleaseFinish(result *AsyncResult) error {
	var err *C.GError
	c := C.g_permission_release_finish(v.native(), result.native(), &err)
	released := gobool(c)
	if !released {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// ImplUpdate is a wrapper around g_permission_impl_update().
func (v *Permission) ImplUpdate(allowed, canAcquire, canRelease bool) {
	C.g_permission_impl_update(v.native(), gbool(allowed), gbool(canAcquire), gbool(canRelease))
}
