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
	ANTIALIAS_DEFAULT  cairo.Antialias = C.CAIRO_ANTIALIAS_DEFAULT
	ANTIALIAS_NONE     cairo.Antialias = C.CAIRO_ANTIALIAS_NONE
	ANTIALIAS_GRAY     cairo.Antialias = C.CAIRO_ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL cairo.Antialias = C.CAIRO_ANTIALIAS_SUBPIXEL
	// ANTIALIAS_FAST     iface.Antialias = C.CAIRO_ANTIALIAS_FAST (since 1.12)
	// ANTIALIAS_GOOD     iface.Antialias = C.CAIRO_ANTIALIAS_GOOD (since 1.12)
	// ANTIALIAS_BEST     iface.Antialias = C.CAIRO_ANTIALIAS_BEST (since 1.12)
)

func marshalAntialias(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.Antialias(c), nil
}
