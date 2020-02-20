package cairo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"
)

// Format is a representation of Cairo's cairo_format_t.
type Format int

const (
	FORMAT_INVALID   Format = C.CAIRO_FORMAT_INVALID
	FORMAT_ARGB32    Format = C.CAIRO_FORMAT_ARGB32
	FORMAT_RGB24     Format = C.CAIRO_FORMAT_RGB24
	FORMAT_A8        Format = C.CAIRO_FORMAT_A8
	FORMAT_A1        Format = C.CAIRO_FORMAT_A1
	FORMAT_RGB16_565 Format = C.CAIRO_FORMAT_RGB16_565
	FORMAT_RGB30     Format = C.CAIRO_FORMAT_RGB30
)

func marshalFormat(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Format(c), nil
}

// FormatStrideForWidth is a wrapper for cairo_format_stride_for_width().
func FormatStrideForWidth(format Format, width int) int {
	c := C.cairo_format_stride_for_width(C.cairo_format_t(format), C.int(width))
	return int(c)
}
