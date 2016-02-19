package cairof

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
	cairo.LINE_JOIN_MITER = C.CAIRO_LINE_JOIN_MITER
	cairo.LINE_JOIN_ROUND = C.CAIRO_LINE_JOIN_ROUND
	cairo.LINE_JOIN_BEVEL = C.CAIRO_LINE_JOIN_BEVEL
}

func marshalLineJoin(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.LineJoin(c), nil
}
