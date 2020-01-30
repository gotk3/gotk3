package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// Permission is a representation of GIO's GPermission.
type Permission struct {
	*Object
}

// Native returns a pointer to the underlying GPermission.
func (v *Permission) Native() *C.GPermission {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGPermission(unsafe.Pointer(v.GObject))
}

func marshalPermission(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPermission(wrapObject(unsafe.Pointer(c))), nil
}

func wrapPermission(obj *Object) *Permission {
	return &Permission{obj}
}

func WrapPermission(ptr unsafe.Pointer) *Permission {
	return &Permission{wrapObject(ptr)}
}

// GetAllowed is a wrapper around g_permission_get_allowed().
func (v *Permission) GetAllowed() bool {
	c := C.g_permission_get_allowed(v.Native())
	return gobool(c)
}

// GetCanAcquire is a wrapper around g_permission_get_can_acquire().
func (v *Permission) GetCanAcquire() bool {
	c := C.g_permission_get_can_acquire(v.Native())
	return gobool(c)
}

// GetCanRelease is a wrapper around g_permission_get_can_release().
func (v *Permission) GetCanRelease() bool {
	c := C.g_permission_get_can_release(v.Native())
	return gobool(c)
}

// Acquire is a wrapper around g_permission_acquire().
// func (v *Permission) Acquire(cancellable *Cancellable) error {
// 	var err *C.GError
// 	c := C.g_permission_acquire(v.Native(), cancellable.Native(), &err)
// 	acquired := gobool(c)
// 	if !acquired {
// 		defer C.g_error_free(err)
// 		return errors.New(C.GoString((*C.char)(err.message)))
// 	}
// 	return nil
// }

// AcquireAsync is a wrapper around g_permission_acquire_async().
// func (v *Permission) AcquireAsync(cancellable *Cancellable, callback AsyncReadyCallback, data uintptr) {
// 	C.g_permission_acquire_async(v.Native(), cancellable.Native(), , C.gpointer(data))
// }

// AcquireFinish is a wrapper around g_permission_acquire_finish().
// func (v *Permission) AcquireFinish(result *AsyncResult) error {
// 	var err *C.GError
// 	c := C.g_permission_acquire_finish(v.Native(), , &err)
// 	acquired := gobool(c)
// 	if !acquired {
// 		defer C.g_error_free(err)
// 		return errors.New(C.GoString((*C.char)(err.message)))
// 	}
// 	return nil
// }

// Release is a wrapper around g_permission_release().
// func (v *Permission) Release(cancellable *Cancellable) error {
// 	var err *C.GError
// 	c := C.g_permission_release(v.Native(), cancellable.Native(), &err)
// 	released := gobool(c)
// 	if !released {
// 		defer C.g_error_free(err)
// 		return errors.New(C.GoString((*C.char)(err.message)))
// 	}
// 	return nil
// }

// ReleaseAsync is a wrapper around g_permission_release_async().
// func (v *Permission) ReleaseAsync(cancellable *Cancellable, callback AsyncReadyCallback, data uintptr) {
// 	C.g_permission_release_async(v.Native(), cancellable.Native(), , C.gpointer(data))
// }

// ReleaseFinish is a wrapper around g_permission_release_finish().
// func (v *Permission) ReleaseFinish(result *AsyncResult) error {
// 	var err *C.GError
// 	c := C.g_permission_release_finish(v.Native(), , &err)
// 	released := gobool(c)
// 	if !released {
// 		defer C.g_error_free(err)
// 		return errors.New(C.GoString((*C.char)(err.message)))
// 	}
// 	return nil
// }

// ImplUpdate is a wrapper around g_permission_impl_update().
func (v *Permission) ImplUpdate(allowed, canAcquire, canRelease bool) {
	C.g_permission_impl_update(v.Native(), gbool(allowed), gbool(canAcquire), gbool(canRelease))
}
