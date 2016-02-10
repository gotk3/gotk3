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
	OPERATOR_CLEAR          iface.Operator = C.CAIRO_OPERATOR_CLEAR
	OPERATOR_SOURCE         iface.Operator = C.CAIRO_OPERATOR_SOURCE
	OPERATOR_OVER           iface.Operator = C.CAIRO_OPERATOR_OVER
	OPERATOR_IN             iface.Operator = C.CAIRO_OPERATOR_IN
	OPERATOR_OUT            iface.Operator = C.CAIRO_OPERATOR_OUT
	OPERATOR_ATOP           iface.Operator = C.CAIRO_OPERATOR_ATOP
	OPERATOR_DEST           iface.Operator = C.CAIRO_OPERATOR_DEST
	OPERATOR_DEST_OVER      iface.Operator = C.CAIRO_OPERATOR_DEST_OVER
	OPERATOR_DEST_IN        iface.Operator = C.CAIRO_OPERATOR_DEST_IN
	OPERATOR_DEST_OUT       iface.Operator = C.CAIRO_OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP      iface.Operator = C.CAIRO_OPERATOR_DEST_ATOP
	OPERATOR_XOR            iface.Operator = C.CAIRO_OPERATOR_XOR
	OPERATOR_ADD            iface.Operator = C.CAIRO_OPERATOR_ADD
	OPERATOR_SATURATE       iface.Operator = C.CAIRO_OPERATOR_SATURATE
	OPERATOR_MULTIPLY       iface.Operator = C.CAIRO_OPERATOR_MULTIPLY
	OPERATOR_SCREEN         iface.Operator = C.CAIRO_OPERATOR_SCREEN
	OPERATOR_OVERLAY        iface.Operator = C.CAIRO_OPERATOR_OVERLAY
	OPERATOR_DARKEN         iface.Operator = C.CAIRO_OPERATOR_DARKEN
	OPERATOR_LIGHTEN        iface.Operator = C.CAIRO_OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE    iface.Operator = C.CAIRO_OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN     iface.Operator = C.CAIRO_OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT     iface.Operator = C.CAIRO_OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT     iface.Operator = C.CAIRO_OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE     iface.Operator = C.CAIRO_OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION      iface.Operator = C.CAIRO_OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE        iface.Operator = C.CAIRO_OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION iface.Operator = C.CAIRO_OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR      iface.Operator = C.CAIRO_OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY iface.Operator = C.CAIRO_OPERATOR_HSL_LUMINOSITY
)

func marshalOperator(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.Operator(c), nil
}
