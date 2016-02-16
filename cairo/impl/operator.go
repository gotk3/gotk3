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

func init() {
	cairo.OPERATOR_CLEAR = C.CAIRO_OPERATOR_CLEAR
	cairo.OPERATOR_SOURCE = C.CAIRO_OPERATOR_SOURCE
	cairo.OPERATOR_OVER = C.CAIRO_OPERATOR_OVER
	cairo.OPERATOR_IN = C.CAIRO_OPERATOR_IN
	cairo.OPERATOR_OUT = C.CAIRO_OPERATOR_OUT
	cairo.OPERATOR_ATOP = C.CAIRO_OPERATOR_ATOP
	cairo.OPERATOR_DEST = C.CAIRO_OPERATOR_DEST
	cairo.OPERATOR_DEST_OVER = C.CAIRO_OPERATOR_DEST_OVER
	cairo.OPERATOR_DEST_IN = C.CAIRO_OPERATOR_DEST_IN
	cairo.OPERATOR_DEST_OUT = C.CAIRO_OPERATOR_DEST_OUT
	cairo.OPERATOR_DEST_ATOP = C.CAIRO_OPERATOR_DEST_ATOP
	cairo.OPERATOR_XOR = C.CAIRO_OPERATOR_XOR
	cairo.OPERATOR_ADD = C.CAIRO_OPERATOR_ADD
	cairo.OPERATOR_SATURATE = C.CAIRO_OPERATOR_SATURATE
	cairo.OPERATOR_MULTIPLY = C.CAIRO_OPERATOR_MULTIPLY
	cairo.OPERATOR_SCREEN = C.CAIRO_OPERATOR_SCREEN
	cairo.OPERATOR_OVERLAY = C.CAIRO_OPERATOR_OVERLAY
	cairo.OPERATOR_DARKEN = C.CAIRO_OPERATOR_DARKEN
	cairo.OPERATOR_LIGHTEN = C.CAIRO_OPERATOR_LIGHTEN
	cairo.OPERATOR_COLOR_DODGE = C.CAIRO_OPERATOR_COLOR_DODGE
	cairo.OPERATOR_COLOR_BURN = C.CAIRO_OPERATOR_COLOR_BURN
	cairo.OPERATOR_HARD_LIGHT = C.CAIRO_OPERATOR_HARD_LIGHT
	cairo.OPERATOR_SOFT_LIGHT = C.CAIRO_OPERATOR_SOFT_LIGHT
	cairo.OPERATOR_DIFFERENCE = C.CAIRO_OPERATOR_DIFFERENCE
	cairo.OPERATOR_EXCLUSION = C.CAIRO_OPERATOR_EXCLUSION
	cairo.OPERATOR_HSL_HUE = C.CAIRO_OPERATOR_HSL_HUE
	cairo.OPERATOR_HSL_SATURATION = C.CAIRO_OPERATOR_HSL_SATURATION
	cairo.OPERATOR_HSL_COLOR = C.CAIRO_OPERATOR_HSL_COLOR
	cairo.OPERATOR_HSL_LUMINOSITY = C.CAIRO_OPERATOR_HSL_LUMINOSITY
}

func marshalOperator(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return cairo.Operator(c), nil
}
