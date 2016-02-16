package impl

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
)

const (
	SURFACE_TYPE_IMAGE          cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_IMAGE
	SURFACE_TYPE_PDF            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_PDF
	SURFACE_TYPE_PS             cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_PS
	SURFACE_TYPE_XLIB           cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_XLIB
	SURFACE_TYPE_XCB            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_XCB
	SURFACE_TYPE_GLITZ          cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_GLITZ
	SURFACE_TYPE_QUARTZ         cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_QUARTZ
	SURFACE_TYPE_WIN32          cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_WIN32
	SURFACE_TYPE_BEOS           cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_BEOS
	SURFACE_TYPE_DIRECTFB       cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_DIRECTFB
	SURFACE_TYPE_SVG            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_SVG
	SURFACE_TYPE_OS2            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_OS2
	SURFACE_TYPE_WIN32_PRINTING cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_WIN32_PRINTING
	SURFACE_TYPE_QUARTZ_IMAGE   cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_QUARTZ_IMAGE
	SURFACE_TYPE_SCRIPT         cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_SCRIPT
	SURFACE_TYPE_QT             cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_QT
	SURFACE_TYPE_RECORDING      cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_RECORDING
	SURFACE_TYPE_VG             cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_VG
	SURFACE_TYPE_GL             cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_GL
	SURFACE_TYPE_DRM            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_DRM
	SURFACE_TYPE_TEE            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_TEE
	SURFACE_TYPE_XML            cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_XML
	SURFACE_TYPE_SKIA           cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_SKIA
	SURFACE_TYPE_SUBSURFACE     cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_SUBSURFACE
	// SURFACE_TYPE_COGL           cairo.SurfaceType = C.CAIRO_SURFACE_TYPE_COGL (since 1.12)
)

func marshalSurfaceType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.SurfaceType(c), nil
}
