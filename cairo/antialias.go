package cairo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"
)

// Antialias is a representation of Cairo's cairo_antialias_t.
type Antialias int

const (
	ANTIALIAS_DEFAULT  Antialias = C.CAIRO_ANTIALIAS_DEFAULT
	ANTIALIAS_NONE     Antialias = C.CAIRO_ANTIALIAS_NONE
	ANTIALIAS_GRAY     Antialias = C.CAIRO_ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL Antialias = C.CAIRO_ANTIALIAS_SUBPIXEL
)

func marshalAntialias(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Antialias(c), nil
}
