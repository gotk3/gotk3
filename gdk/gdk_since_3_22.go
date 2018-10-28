// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20
// Supports building with gtk 3.22+

// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk_since_3_22.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GdkDisplay
 */

// GetNMonitors is a wrapper around gdk_display_get_n_monitors().
func (v *Display) GetNMonitors() int {
	c := C.gdk_display_get_n_monitors(v.native())
	return int(c)
}

// GetPrimaryMonitor is a wrapper around gdk_display_get_primary_monitor().
func (v *Display) GetPrimaryMonitor() (*Monitor, error) {
	c := C.gdk_display_get_primary_monitor(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Monitor{glib.Take(unsafe.Pointer(c))}, nil
}

/*
 * GdkMonitor
 */

// Monitor is a representation of GDK's GdkMonitor.
type Monitor struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkMonitor.
func (v *Monitor) native() *C.GdkMonitor {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkMonitor(p)
}

// Native returns a pointer to the underlying GdkMonitor.
func (v *Monitor) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalMonitor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Monitor{obj}, nil
}

func toMonitor(s *C.GdkMonitor) (*Monitor, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
	return &Monitor{obj}, nil
}

// GetGeometry is a wrapper around gdk_monitor_get_geometry().
func (v *Monitor) GetGeometry() *Rectangle {
	var rect C.GdkRectangle

	C.gdk_monitor_get_geometry(v.native(), &rect)

	return WrapRectangle(uintptr(unsafe.Pointer(&rect)))
}
