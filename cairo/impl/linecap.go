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
	LINE_CAP_BUTT   cairo.LineCap = C.CAIRO_LINE_CAP_BUTT
	LINE_CAP_ROUND  cairo.LineCap = C.CAIRO_LINE_CAP_ROUND
	LINE_CAP_SQUARE cairo.LineCap = C.CAIRO_LINE_CAP_SQUARE
)

func marshalLineCap(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.LineCap(c), nil
}
