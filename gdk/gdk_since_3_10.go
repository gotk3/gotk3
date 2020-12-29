// +build !gtk_3_6,!gtk_3_8
// Supports building with gtk 3.10+

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
)

// TODO:
// gdk_device_get_position_double().

// GetScaleFactor is a wrapper around gdk_window_get_scale_factor().
func (v *Window) GetScaleFactor() int {
	return int(C.gdk_window_get_scale_factor(v.native()))
}

// CreateSimilarImageSurface is a wrapper around gdk_window_create_similar_image_surface().
func (v *Window) CreateSimilarImageSurface(format cairo.Format, w, h, scale int) (*cairo.Surface, error) {
	surface := C.gdk_window_create_similar_image_surface(v.native(), C.cairo_format_t(format), C.gint(w), C.gint(h), C.gint(scale))

	status := cairo.Status(C.cairo_surface_status(surface))
	if status != cairo.STATUS_SUCCESS {
		return nil, cairo.ErrorStatus(status)
	}

	return cairo.NewSurface(uintptr(unsafe.Pointer(surface)), false), nil
}

// CairoSurfaceCreateFromPixbuf is a wrapper around gdk_cairo_surface_create_from_pixbuf().
func CairoSurfaceCreateFromPixbuf(pixbuf *Pixbuf, scale int, window *Window) (*cairo.Surface, error) {
	v := C.gdk_cairo_surface_create_from_pixbuf(pixbuf.native(), C.gint(scale), window.native())

	status := cairo.Status(C.cairo_surface_status(v))
	if status != cairo.STATUS_SUCCESS {
		return nil, cairo.ErrorStatus(status)
	}

	surface := cairo.WrapSurface(uintptr(unsafe.Pointer(v)))
	// Keep pixbuf alive.
	runtime.SetFinalizer(surface, func(surface *cairo.Surface) {
		runtime.KeepAlive(pixbuf)
		surface.Close()
	})

	return surface, nil
}
