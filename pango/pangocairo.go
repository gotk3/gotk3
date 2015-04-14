/*
 * Copyright (c) 2015- terrak <terrak1975@gmail.com>
 *
 * This file originated from: http://www.terrak.net/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package pango

// #cgo pkg-config: pango pangocairo
// #include <pango/pango.h>
// #include <pango/pangocairo.h>
// #include "pango.go.h"
import "C"
import (
	//	"github.com/terrak/gotk3/glib"
	"github.com/terrak/gotk3/cairo"
	"unsafe"
)

func init() {
	//	tm := []glib.TypeMarshaler{
	//		// Enums
	//		{glib.Type(C.pango_alignement_get_type()), marshalAlignment},
	//		{glib.Type(C.pango_ellipsize_mode_get_type()), marshalEllipsizeMode},
	//		{glib.Type(C.pango_wrap_mode_get_type()), marshalWrapMode},
	//	}
	//	glib.RegisterGValueMarshalers(tm)
}

/* Convenience
 */
//PangoContext *pango_cairo_create_context (cairo_t   *cr);
func (cr *cairo.Context) CairoCreateContext() *Context {
	c := C.pango_cairo_create_context(cr.Native())
	context := new(Context)
	context.pangoContext = (*C.PangoContext)(c)
	return context
}

//PangoLayout *pango_cairo_create_layout (cairo_t     *cr);
func (cr *cairo.Context) CairoCreateLayout() *Layout {
	c := C.pango_cairo_create_layout(cr.Native())
	layout := new(Layout)
	layout.pangoLayout = (*C.PangoLayout)(c)
	return layout
}

//void         pango_cairo_update_layout (cairo_t     *cr,
//					PangoLayout *layout);
func (cr *cairo.Context) CairoUpdateLayout(v *Layout) {
	C.pango_cairo_update_layout(cr.Native(), v.native())
}

/*
 * Rendering
 */
//void pango_cairo_show_glyph_string (cairo_t          *cr,
//				    PangoFont        *font,
//				    PangoGlyphString *glyphs);
func (cr *cairo.Context) CairoShowGlyphString(font *Font, glyphs *GlyphString) {
	C.pango_cairo_show_glyph_string(cr.Native(), font.native(), glyphs.native())
}

//void pango_cairo_show_glyph_item   (cairo_t          *cr,
//				    const char       *text,
//				    PangoGlyphItem   *glyph_item);
func (cr *cairo.Context) CairoShowGlyphItem(text string, glyph_item *GlyphItem) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_cairo_show_glyph_item(cr.Native(), (*C.char)(cstr), glyph_item.native() )
}

//void pango_cairo_show_layout_line  (cairo_t          *cr,
//				    PangoLayoutLine  *line);
func (cr *cairo.Context)CairoShowLayoutLine(line *LayoutLine){
	C.pango_cairo_show_layout_line(cr.Native(), line.native())
}

//void pango_cairo_show_layout       (cairo_t          *cr,
//				    PangoLayout      *layout);
func (cr *cairo.Context)CairoShowLayout(layout *Layout){
	C.pango_cairo_show_layout(cr.Native(), layout.native())
}


//void pango_cairo_show_error_underline (cairo_t       *cr,
//				       double         x,
//				       double         y,
//				       double         width,
//				       double         height);

/*
 * Rendering to a path
 */

//void pango_cairo_glyph_string_path (cairo_t          *cr,
//				    PangoFont        *font,
//				    PangoGlyphString *glyphs);
func (cr *cairo.Context) CairoGlyphStringPath(font *Font, glyphs *GlyphString) {
	C.pango_cairo_glyph_string_path(cr.Native(), font.native(), glyphs.native())
}

//void pango_cairo_layout_line_path  (cairo_t          *cr,
//				    PangoLayoutLine  *line);
func (cr *cairo.Context)CairoLayoutLinePath(line *LayoutLine){
	C.pango_cairo_layout_line_path(cr.Native(), line.native())
}

//void pango_cairo_layout_path       (cairo_t          *cr,
//				    PangoLayout      *layout);
func (cr *cairo.Context)CairoLayoutPath(layout *Layout){
	C.pango_cairo_layout_path(cr.Native(), layout.native())
}

//void pango_cairo_error_underline_path (cairo_t       *cr,
//				       double         x,
//				       double         y,
//				       double         width,
//				       double         height);
func (cr *cairo.Context)CairoErrorUnderlinePath(x,y,width,height float64){
	C.pango_cairo_error_underline_path(cr.native(), C.double(x), C.double(y), C.double(width), C.double(height))
}

