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
	FILL_RULE_WINDING  iface.FillRule = C.CAIRO_FILL_RULE_WINDING
	FILL_RULE_EVEN_ODD iface.FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)

func marshalFillRule(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.FillRule(c), nil
}
