// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

package gdk

// #include <gdk/gdk.h>
// #include "glarea_since_3_16.go.h"
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

type (
	// MajorVersion is a representation of OpenGL major version.
	MajorVersion int
	// MinorVersion is a representation of OpenGL minor version.
	MinorVersion int
)

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

// GetUseES is a wrapper around gdk_gl_context_get_use_es().
func (v *GLContext) GetUseES() bool {
	return gobool(C.gdk_gl_context_get_use_es(v.native()))
}

// SetUseES is a wrapper around gdk_gl_context_set_use_es().
func (v *GLContext) SetUseES(es int) {
	C.gdk_gl_context_set_use_es(v.native(), (C.int)(es))
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

// IsLegacy is a wrapper around gdk_gl_context_is_legacy().
func (v *GLContext) IsLegacy() bool {
	return gobool(C.gdk_gl_context_is_legacy(v.native()))
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
