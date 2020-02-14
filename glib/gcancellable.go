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

// Cancellable is a representation of GIO's GCancellable.
type Cancellable struct {
	*Object
}

// native returns a pointer to the underlying GCancellable.
func (v *Cancellable) native() *C.GCancellable {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toCancellable(unsafe.Pointer(v.GObject))
}

func marshalCancellable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapCancellable(wrapObject(unsafe.Pointer(c))), nil
}

func wrapCancellable(obj *Object) *Cancellable {
	return &Cancellable{obj}
}

// CancellableNew is a wrapper around g_cancellable_new().
func CancellableNew() (*Cancellable, error) {
	c := C.g_cancellable_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapCancellable(wrapObject(unsafe.Pointer(c))), nil
}

// IsCancelled is a wrapper around g_cancellable_is_cancelled().
func (v *Cancellable) IsCancelled() bool {
	c := C.g_cancellable_is_cancelled(v.native())
	return gobool(c)
}

// SetErrorIfCancelled is a wrapper around g_cancellable_set_error_if_cancelled().
func (v *Cancellable) SetErrorIfCancelled() error {
	var err *C.GError
	c := C.g_cancellable_set_error_if_cancelled(v.native(), &err)
	cancelled := gobool(c)
	if cancelled {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// GetFD is a wrapper around g_cancellable_get_fd().
func (v *Cancellable) GetFD() int {
	c := C.g_cancellable_get_fd(v.native())
	return int(c)
}

// MakePollFD is a wrapper around g_cancellable_make_pollfd().
// func (v *Cancellable) MakePollFD(pollFD *PollFD) bool {
// 	c := C.g_cancellable_make_pollfd(v.native(), )
// 	return gobool(c)
// }

// ReleaseFD is a wrapper around g_cancellable_release_fd().
func (v *Cancellable) ReleaseFD() {
	C.g_cancellable_release_fd(v.native())
}

// SourceNew is a wrapper around g_cancellable_source_new().
func (v *Cancellable) SourceNew() *Source {
	c := C.g_cancellable_source_new(v.native())
	return wrapSource(c)
}
