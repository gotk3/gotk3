// +build !pango_1_10,!pango_1_12,!pango_1_14

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
