// +build !cairo_1_9,!cairo_1_10,!cairo_1_11,!cairo_1_12,!cairo_1_13,!cairo_1_14,!cairo_1_15

package cairo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"
)

// GetVariations is a wrapper around cairo_font_options_get_variations().
func (o *FontOptions) GetVariations() string {
	return C.GoString(C.cairo_font_options_get_variations(o.native))
}

// SetVariations is a wrapper around cairo_font_options_set_variations().
func (o *FontOptions) SetVariations(variations string) {
	var cvariations *C.char
	if variations != "" {
		cvariations = C.CString(variations)
		// Cairo will call strdup on its own.
		defer C.free(unsafe.Pointer(cvariations))
	}

	C.cairo_font_options_set_variations(o.native, cvariations)
}
