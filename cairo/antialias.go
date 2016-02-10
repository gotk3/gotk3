package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"
)

const (
	ANTIALIAS_DEFAULT  Antialias = C.CAIRO_ANTIALIAS_DEFAULT
	ANTIALIAS_NONE     Antialias = C.CAIRO_ANTIALIAS_NONE
	ANTIALIAS_GRAY     Antialias = C.CAIRO_ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL Antialias = C.CAIRO_ANTIALIAS_SUBPIXEL
	// ANTIALIAS_FAST     Antialias = C.CAIRO_ANTIALIAS_FAST (since 1.12)
	// ANTIALIAS_GOOD     Antialias = C.CAIRO_ANTIALIAS_GOOD (since 1.12)
	// ANTIALIAS_BEST     Antialias = C.CAIRO_ANTIALIAS_BEST (since 1.12)
)

func marshalAntialias(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Antialias(c), nil
}
