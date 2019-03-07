//+build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

package gtk

// #include <gtk/gtk.h>
// #include "glarea_since_3_16.go.h"
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
	return &GLArea{Widget{glib.InitiallyUnowned{obj}}}
}

func WidgetToGLArea(widget *Widget) (interface{}, error) {
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

// GetUseES is a wrapper around gtk_gl_area_get_use_es().
func (v *GLArea) GetUseES() bool {
	return gobool(C.gtk_gl_area_get_use_es(v.native()))
}

// SetUseES is a wrapper around gtk_gl_area_set_use_es().
func (v *GLArea) SetUseES(es bool) {
	C.gtk_gl_area_set_use_es(v.native(), gbool(es))
}

type (
	// MajorVersion is a representation of OpenGL major version.
	MajorVersion int
	// MinorVersion is a representation of OpenGL minor version.
	MinorVersion int
)

// GetRequiredVersion is a wrapper around gtk_gl_area_get_required_version().
func (v *GLArea) GetRequiredVersion() (MajorVersion, MinorVersion) {
	var major, minor int
	C.gtk_gl_area_get_required_version(v.native(),
		(*C.int)(unsafe.Pointer(&major)), (*C.int)(unsafe.Pointer(&minor)))

	return MajorVersion(major), MinorVersion(minor)
}

// SetRequiredVersion is a wrapper around gtk_gl_area_set_required_version().
func (v *GLArea) SetRequiredVersion(major, minor int) {
	C.gtk_gl_area_set_required_version(v.native(), (C.int)(major), (C.int)(minor))
}

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

// SetAutoRender is a wrapper around gtk_gl_area_set_has_stencil_buffer().
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

// GError* gtk_gl_area_get_error (GtkGLArea *area);
func (v *GLArea) GetError() error {
	var err *C.GError = nil
	err = C.gtk_gl_area_get_error(v.native())
	if err != nil {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// void gtk_gl_area_set_error (GtkGLArea *area, const GError *error);
