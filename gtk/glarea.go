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

package gtk

// #include <gtk/gtk.h>
// #include "gtk_since_3_16.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_gl_area_get_type()), marshalGLArea},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkGLArea"] = wrapGLArea
}

/*
 * GtkGLArea
 */

// GLArea is a representation of GTK's GtkGLArea.
type GLArea struct {
	Widget
}

// native returns a pointer to the underlying GtkGLArea.
func (v *GLArea) native() *C.GtkGLArea {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkGLArea(p)
}

func marshalGLArea(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapGLArea(obj), nil
}

func wrapGLArea(obj *glib.Object) *GLArea {
	if obj == nil {
		return nil
	}

	return &GLArea{Widget{glib.InitiallyUnowned{obj}}}
}

func WidgetToGLArea(widget *Widget) (*GLArea, error) {
	obj := glib.Take(unsafe.Pointer(widget.GObject))
	return wrapGLArea(obj), nil
}

// GLAreaNew is a wrapper around gtk_gl_area_new().
func GLAreaNew() (*GLArea, error) {
	c := C.gtk_gl_area_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))

	return wrapGLArea(obj), nil
}

// MajorVersion is a representation of OpenGL major version.
type MajorVersion int

// MinorVersion is a representation of OpenGL minor version.
type MinorVersion int

// GetRequiredVersion is a wrapper around gtk_gl_area_get_required_version().
func (v *GLArea) GetRequiredVersion() (MajorVersion, MinorVersion) {
	var major, minor int
	C.gtk_gl_area_get_required_version(v.native(),
		(*C.gint)(unsafe.Pointer(&major)), (*C.gint)(unsafe.Pointer(&minor)))

	return MajorVersion(major), MinorVersion(minor)
}

// SetRequiredVersion is a wrapper around gtk_gl_area_set_required_version().
func (v *GLArea) SetRequiredVersion(major, minor int) {
	C.gtk_gl_area_set_required_version(v.native(), (C.gint)(major), (C.gint)(minor))
}

// TODO:
// gtk_gl_area_set_has_alpha().
// gtk_gl_area_get_has_alpha().

// HasDepthBuffer is a wrapper around gtk_gl_area_get_has_depth_buffer().
func (v *GLArea) HasDepthBuffer() bool {
	return gobool(C.gtk_gl_area_get_has_depth_buffer(v.native()))
}

// SetHasDepthBuffer is a wrapper around gtk_gl_area_set_has_depth_buffer().
func (v *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	C.gtk_gl_area_set_has_depth_buffer(v.native(), gbool(hasDepthBuffer))
}

// HasStencilBuffer is a wrapper around gtk_gl_area_get_has_stencil_buffer().
func (v *GLArea) HasStencilBuffer() bool {
	return gobool(C.gtk_gl_area_get_has_stencil_buffer(v.native()))
}

// SetHasStencilBuffer is a wrapper around gtk_gl_area_set_has_stencil_buffer().
func (v *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	C.gtk_gl_area_set_has_stencil_buffer(v.native(), gbool(hasStencilBuffer))
}

// GetAutoRender is a wrapper around gtk_gl_area_get_auto_render().
func (v *GLArea) GetAutoRender() bool {
	return gobool(C.gtk_gl_area_get_auto_render(v.native()))
}

// SetAutoRender is a wrapper around gtk_gl_area_set_auto_render().
func (v *GLArea) SetAutoRender(autoRender bool) {
	C.gtk_gl_area_set_auto_render(v.native(), gbool(autoRender))
}

// QueueRender is a wrapper around gtk_gl_area_queue_render().
func (v *GLArea) QueueRender() {
	C.gtk_gl_area_queue_render(v.native())
}

// GetContext is a wrapper around gtk_gl_area_get_context().
func (v *GLArea) GetContext() (*gdk.GLContext, error) {
	c := C.gtk_gl_area_get_context(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &(gdk.GLContext{obj}), nil
}

// MakeCurrent is a wrapper around gtk_gl_area_make_current().
func (v *GLArea) MakeCurrent() {
	C.gtk_gl_area_make_current(v.native())
}

// AttachBuffers is a wrapper around gtk_gl_area_attach_buffers().
func (v *GLArea) AttachBuffers() {
	C.gtk_gl_area_attach_buffers(v.native())
}

// GetError is a wrapper around gtk_gl_area_get_error().
func (v *GLArea) GetError() error {
	var err *C.GError = nil
	err = C.gtk_gl_area_get_error(v.native())
	if err != nil {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// SetError is a wrapper around gtk_gl_area_set_error().
func (v *GLArea) SetError(domain glib.Quark, code int, err error) {
	cstr := (*C.gchar)(C.CString(err.Error()))
	defer C.free(unsafe.Pointer(cstr))

	gerr := C.g_error_new_literal(C.GQuark(domain), C.gint(code), cstr)
	defer C.g_error_free(gerr)

	C.gtk_gl_area_set_error(v.native(), gerr)
}