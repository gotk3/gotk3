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
	LINE_CAP_BUTT   iface.LineCap = C.CAIRO_LINE_CAP_BUTT
	LINE_CAP_ROUND  iface.LineCap = C.CAIRO_LINE_CAP_ROUND
	LINE_CAP_SQUARE iface.LineCap = C.CAIRO_LINE_CAP_SQUARE
)

func marshalLineCap(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.LineCap(c), nil
}
