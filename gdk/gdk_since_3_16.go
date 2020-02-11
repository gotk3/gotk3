// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14
// Supports building with gtk 3.16+

/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gdk

// #include <gdk/gdk.h>
// #include "gdk_since_3_16.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gdk_gl_context_get_type()), marshalGLContext},
	}

	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

const (
	GRAB_FAILED GrabStatus = C.GDK_GRAB_FAILED
)

/*
 * GdkDevice
 */

// TODO:
// gdk_device_get_vendor_id().
// gdk_device_get_product_id().

/*
 * GdkGLContext
 */

// GLContext is a representation of GDK's GdkGLContext.
type GLContext struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkGLContext.
func (v *GLContext) native() *C.GdkGLContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkGLContext(p)
}

// Native returns a pointer to the underlying GdkGLContext.
func (v *GLContext) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalGLContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &GLContext{obj}, nil
}

// GetDisplay is a wrapper around gdk_gl_context_get_display().
func (v *GLContext) GetDisplay() (*Display, error) {
	c := C.gdk_gl_context_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// GetWindow is a wrapper around gdk_gl_context_get_window().
func (v *GLContext) GetSurface() (*Window, error) {
	c := C.gdk_gl_context_get_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Window{glib.Take(unsafe.Pointer(c))}, nil
}

// GetSharedContext is a wrapper around gdk_gl_context_get_shared_context().
func (v *GLContext) GetSharedContext() (*GLContext, error) {
	c := C.gdk_gl_context_get_shared_context(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &GLContext{glib.Take(unsafe.Pointer(c))}, nil
}

// MajorVersion is a representation of OpenGL major version.
type MajorVersion int

// MinorVersion is a representation of OpenGL minor version.
type MinorVersion int

// GetVersion is a wrapper around gdk_gl_context_get_version().
func (v *GLContext) GetVersion() (MajorVersion, MinorVersion) {
	var major, minor int
	C.gdk_gl_context_get_version(v.native(),
		(*C.int)(unsafe.Pointer(&major)), (*C.int)(unsafe.Pointer(&minor)))

	return MajorVersion(major), MinorVersion(minor)
}

// GetRequiredVersion is a wrapper around gdk_gl_context_get_required_version().
func (v *GLContext) GetRequiredVersion() (MajorVersion, MinorVersion) {
	var major, minor int
	C.gdk_gl_context_get_required_version(v.native(),
		(*C.int)(unsafe.Pointer(&major)), (*C.int)(unsafe.Pointer(&minor)))

	return MajorVersion(major), MinorVersion(minor)
}

// SetRequiredVersion is a wrapper around gdk_gl_context_set_required_version().
func (v *GLContext) SetRequiredVersion(major, minor int) {
	C.gdk_gl_context_set_required_version(v.native(), (C.int)(major), (C.int)(minor))
}

// GetDebugEnabled is a wrapper around gdk_gl_context_get_debug_enabled().
func (v *GLContext) GetDebugEnabled() bool {
	return gobool(C.gdk_gl_context_get_debug_enabled(v.native()))
}

// SetDebugEnabled is a wrapper around gdk_gl_context_set_debug_enabled().
func (v *GLContext) SetDebugEnabled(enabled bool) {
	C.gdk_gl_context_set_debug_enabled(v.native(), gbool(enabled))
}

// GetForwardCompatible is a wrapper around gdk_gl_context_get_forward_compatible().
func (v *GLContext) GetForwardCompatible() bool {
	return gobool(C.gdk_gl_context_get_forward_compatible(v.native()))
}

// SetForwardCompatible is a wrapper around gdk_gl_context_set_forward_compatible().
func (v *GLContext) SetForwardCompatible(compatible bool) {
	C.gdk_gl_context_set_forward_compatible(v.native(), gbool(compatible))
}

// Realize is a wrapper around gdk_gl_context_realize().
func (v *GLContext) Realize() (bool, error) {
	var err *C.GError
	r := gobool(C.gdk_gl_context_realize(v.native(), &err))
	if !r {
		defer C.g_error_free(err)
		return r, errors.New(C.GoString((*C.char)(err.message)))
	}

	return r, nil
}

// MakeCurrent is a wrapper around gdk_gl_context_make_current().
func (v *GLContext) MakeCurrent() {
	C.gdk_gl_context_make_current(v.native())
}

// GetCurrent is a wrapper around gdk_gl_context_get_current().
func GetCurrent() (*GLContext, error) {
	c := C.gdk_gl_context_get_current()
	if c == nil {
		return nil, nilPtrErr
	}

	return &GLContext{glib.Take(unsafe.Pointer(c))}, nil
}

// ClearCurrent is a wrapper around gdk_gl_context_clear_current().
func ClearCurrent() {
	C.gdk_gl_context_clear_current()
}
