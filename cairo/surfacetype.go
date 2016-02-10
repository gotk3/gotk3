package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/cairo/iface"
)

const (
	SURFACE_TYPE_IMAGE          iface.SurfaceType = C.CAIRO_SURFACE_TYPE_IMAGE
	SURFACE_TYPE_PDF            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_PDF
	SURFACE_TYPE_PS             iface.SurfaceType = C.CAIRO_SURFACE_TYPE_PS
	SURFACE_TYPE_XLIB           iface.SurfaceType = C.CAIRO_SURFACE_TYPE_XLIB
	SURFACE_TYPE_XCB            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_XCB
	SURFACE_TYPE_GLITZ          iface.SurfaceType = C.CAIRO_SURFACE_TYPE_GLITZ
	SURFACE_TYPE_QUARTZ         iface.SurfaceType = C.CAIRO_SURFACE_TYPE_QUARTZ
	SURFACE_TYPE_WIN32          iface.SurfaceType = C.CAIRO_SURFACE_TYPE_WIN32
	SURFACE_TYPE_BEOS           iface.SurfaceType = C.CAIRO_SURFACE_TYPE_BEOS
	SURFACE_TYPE_DIRECTFB       iface.SurfaceType = C.CAIRO_SURFACE_TYPE_DIRECTFB
	SURFACE_TYPE_SVG            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_SVG
	SURFACE_TYPE_OS2            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_OS2
	SURFACE_TYPE_WIN32_PRINTING iface.SurfaceType = C.CAIRO_SURFACE_TYPE_WIN32_PRINTING
	SURFACE_TYPE_QUARTZ_IMAGE   iface.SurfaceType = C.CAIRO_SURFACE_TYPE_QUARTZ_IMAGE
	SURFACE_TYPE_SCRIPT         iface.SurfaceType = C.CAIRO_SURFACE_TYPE_SCRIPT
	SURFACE_TYPE_QT             iface.SurfaceType = C.CAIRO_SURFACE_TYPE_QT
	SURFACE_TYPE_RECORDING      iface.SurfaceType = C.CAIRO_SURFACE_TYPE_RECORDING
	SURFACE_TYPE_VG             iface.SurfaceType = C.CAIRO_SURFACE_TYPE_VG
	SURFACE_TYPE_GL             iface.SurfaceType = C.CAIRO_SURFACE_TYPE_GL
	SURFACE_TYPE_DRM            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_DRM
	SURFACE_TYPE_TEE            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_TEE
	SURFACE_TYPE_XML            iface.SurfaceType = C.CAIRO_SURFACE_TYPE_XML
	SURFACE_TYPE_SKIA           iface.SurfaceType = C.CAIRO_SURFACE_TYPE_SKIA
	SURFACE_TYPE_SUBSURFACE     iface.SurfaceType = C.CAIRO_SURFACE_TYPE_SUBSURFACE
	// SURFACE_TYPE_COGL           iface.SurfaceType = C.CAIRO_SURFACE_TYPE_COGL (since 1.12)
)

func marshalSurfaceType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.SurfaceType(c), nil
}
