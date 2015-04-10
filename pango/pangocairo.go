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

// #cgo pkg-config: pango
// #include <pango/pango.h>
// #include "pango.go.h"
import "C"
import (
//	"github.com/terrak/gotk3/glib"
//	"github.com/terrak/gotk3/cairo"
//	"unsafe"
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
//PangoLayout *pango_cairo_create_layout (cairo_t     *cr);
//void         pango_cairo_update_layout (cairo_t     *cr,
//					PangoLayout *layout);


/*
 * Rendering
 */
//void pango_cairo_show_glyph_string (cairo_t          *cr,
//				    PangoFont        *font,
//				    PangoGlyphString *glyphs);
//void pango_cairo_show_glyph_item   (cairo_t          *cr,
//				    const char       *text,
//				    PangoGlyphItem   *glyph_item);
//void pango_cairo_show_layout_line  (cairo_t          *cr,
//				    PangoLayoutLine  *line);
//void pango_cairo_show_layout       (cairo_t          *cr,
//				    PangoLayout      *layout);
//
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
//void pango_cairo_layout_line_path  (cairo_t          *cr,
//				    PangoLayoutLine  *line);
//void pango_cairo_layout_path       (cairo_t          *cr,
//				    PangoLayout      *layout);
//
//void pango_cairo_error_underline_path (cairo_t       *cr,
//				       double         x,
//				       double         y,
//				       double         width,
//				       double         height);
