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
	OPERATOR_CLEAR          cairo.Operator = C.CAIRO_OPERATOR_CLEAR
	OPERATOR_SOURCE         cairo.Operator = C.CAIRO_OPERATOR_SOURCE
	OPERATOR_OVER           cairo.Operator = C.CAIRO_OPERATOR_OVER
	OPERATOR_IN             cairo.Operator = C.CAIRO_OPERATOR_IN
	OPERATOR_OUT            cairo.Operator = C.CAIRO_OPERATOR_OUT
	OPERATOR_ATOP           cairo.Operator = C.CAIRO_OPERATOR_ATOP
	OPERATOR_DEST           cairo.Operator = C.CAIRO_OPERATOR_DEST
	OPERATOR_DEST_OVER      cairo.Operator = C.CAIRO_OPERATOR_DEST_OVER
	OPERATOR_DEST_IN        cairo.Operator = C.CAIRO_OPERATOR_DEST_IN
	OPERATOR_DEST_OUT       cairo.Operator = C.CAIRO_OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP      cairo.Operator = C.CAIRO_OPERATOR_DEST_ATOP
	OPERATOR_XOR            cairo.Operator = C.CAIRO_OPERATOR_XOR
	OPERATOR_ADD            cairo.Operator = C.CAIRO_OPERATOR_ADD
	OPERATOR_SATURATE       cairo.Operator = C.CAIRO_OPERATOR_SATURATE
	OPERATOR_MULTIPLY       cairo.Operator = C.CAIRO_OPERATOR_MULTIPLY
	OPERATOR_SCREEN         cairo.Operator = C.CAIRO_OPERATOR_SCREEN
	OPERATOR_OVERLAY        cairo.Operator = C.CAIRO_OPERATOR_OVERLAY
	OPERATOR_DARKEN         cairo.Operator = C.CAIRO_OPERATOR_DARKEN
	OPERATOR_LIGHTEN        cairo.Operator = C.CAIRO_OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE    cairo.Operator = C.CAIRO_OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN     cairo.Operator = C.CAIRO_OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT     cairo.Operator = C.CAIRO_OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT     cairo.Operator = C.CAIRO_OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE     cairo.Operator = C.CAIRO_OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION      cairo.Operator = C.CAIRO_OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE        cairo.Operator = C.CAIRO_OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION cairo.Operator = C.CAIRO_OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR      cairo.Operator = C.CAIRO_OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY cairo.Operator = C.CAIRO_OPERATOR_HSL_LUMINOSITY
)

func marshalOperator(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.Operator(c), nil
}
