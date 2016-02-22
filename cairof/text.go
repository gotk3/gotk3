package cairof

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
	cairo.FONT_SLANT_NORMAL = C.CAIRO_FONT_SLANT_NORMAL
	cairo.FONT_SLANT_ITALIC = C.CAIRO_FONT_SLANT_ITALIC
	cairo.FONT_SLANT_OBLIQUE = C.CAIRO_FONT_SLANT_OBLIQUE

	cairo.FONT_WEIGHT_NORMAL = C.CAIRO_FONT_WEIGHT_NORMAL
	cairo.FONT_WEIGHT_BOLD = C.CAIRO_FONT_WEIGHT_BOLD
}

func (v *Context) SelectFontFace(family string, slant cairo.FontSlant, weight cairo.FontWeight) {
	cstr := C.CString(family)
	defer C.free(unsafe.Pointer(cstr))
	C.cairo_select_font_face(v.native(), (*C.char)(cstr), C.cairo_font_slant_t(slant), C.cairo_font_weight_t(weight))
}

func (v *Context) SetFontSize(size float64) {
	C.cairo_set_font_size(v.native(), C.double(size))
}

// TODO: cairo_set_font_matrix

// TODO: cairo_get_font_matrix

// TODO: cairo_set_font_options

// TODO: cairo_get_font_options

// TODO: cairo_set_font_face

// TODO: cairo_get_font_face

// TODO: cairo_set_scaled_font

// TODO: cairo_get_scaled_font

func (v *Context) ShowText(utf8 string) {
	cstr := C.CString(utf8)
	defer C.free(unsafe.Pointer(cstr))
	C.cairo_show_text(v.native(), (*C.char)(cstr))
}

// TODO: cairo_show_glyphs

// TODO: cairo_show_text_glyphs

func (v *Context) FontExtents() cairo.FontExtents {
	var extents C.cairo_font_extents_t
	C.cairo_font_extents(v.native(), &extents)
	return cairo.FontExtents{
		Ascent:      float64(extents.ascent),
		Descent:     float64(extents.descent),
		Height:      float64(extents.height),
		MaxXAdvance: float64(extents.max_x_advance),
		MaxYAdvance: float64(extents.max_y_advance),
	}
}

func (v *Context) TextExtents(utf8 string) cairo.TextExtents {
	cstr := C.CString(utf8)
	defer C.free(unsafe.Pointer(cstr))
	var extents C.cairo_text_extents_t
	C.cairo_text_extents(v.native(), (*C.char)(cstr), &extents)
	return cairo.TextExtents{
		XBearing: float64(extents.x_bearing),
		YBearing: float64(extents.y_bearing),
		Width:    float64(extents.width),
		Height:   float64(extents.height),
		XAdvance: float64(extents.x_advance),
		YAdvance: float64(extents.y_advance),
	}
}

// TODO: cairo_glyph_extents

// TODO: cairo_toy_font_face_create

// TODO: cairo_toy_font_face_get_family

// TODO: cairo_toy_font_face_get_slant

// TODO: cairo_toy_font_face_get_weight

// TODO: cairo_glyph_allocate

// TODO: cairo_glyph_free

// TODO: cairo_text_cluster_allocate

// TODO: cairo_text_cluster_free
