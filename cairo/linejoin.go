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
	LINE_JOIN_MITER iface.LineJoin = C.CAIRO_LINE_JOIN_MITER
	LINE_JOIN_ROUND iface.LineJoin = C.CAIRO_LINE_JOIN_ROUND
	LINE_JOIN_BEVEL iface.LineJoin = C.CAIRO_LINE_JOIN_BEVEL
)

func marshalLineJoin(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.LineJoin(c), nil
}
