// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18
// Supports building with gtk 3.20+

package gdk

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// #include <gdk/gdk.h>
// #include "gdk_since_3_20.go.h"
import "C"

/*
 * GdkGLContext
 */

// IsLegacy is a wrapper around gdk_gl_context_is_legacy().
func (v *GLContext) IsLegacy() bool {
	return gobool(C.gdk_gl_context_is_legacy(v.native()))
}

/*
 * GdkDisplay
 */

func (v *Display) GetDefaultSeat() (*Seat, error) {
	return toSeat(C.gdk_display_get_default_seat(v.native()))
}

// gdk_display_list_seats().

/*
 * GdkDevice
 */

// TODO:
// gdk_device_get_axes().
// gdk_device_get_seat().

/*
 * GdkSeat
 */

type Seat struct {
	*glib.Object
}

func (v *Seat) native() *C.GdkSeat {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkSeat(p)
}

// Native returns a pointer to the underlying GdkCursor.
func (v *Seat) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalSeat(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Seat{obj}, nil
}

func toSeat(s *C.GdkSeat) (*Seat, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
	return &Seat{obj}, nil
}

func (v *Seat) GetPointer() (*Device, error) {
	return toDevice(C.gdk_seat_get_pointer(v.native()))
}
