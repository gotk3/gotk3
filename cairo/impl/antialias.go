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

func init() {
	cairo.ANTIALIAS_DEFAULT = C.CAIRO_ANTIALIAS_DEFAULT
	cairo.ANTIALIAS_NONE = C.CAIRO_ANTIALIAS_NONE
	cairo.ANTIALIAS_GRAY = C.CAIRO_ANTIALIAS_GRAY
	cairo.ANTIALIAS_SUBPIXEL = C.CAIRO_ANTIALIAS_SUBPIXEL
	// ANTIALIAS_FAST      = C.CAIRO_ANTIALIAS_FAST (since 1.12)
	// ANTIALIAS_GOOD      = C.CAIRO_ANTIALIAS_GOOD (since 1.12)
	// ANTIALIAS_BEST      = C.CAIRO_ANTIALIAS_BEST (since 1.12)
}

func marshalAntialias(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.Antialias(c), nil
}
