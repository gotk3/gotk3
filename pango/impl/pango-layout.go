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

package impl

// #cgo pkg-config: pango
// #include <pango/pango.h>
// #include "pango.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/pango"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		// Enums
		{glib.Type(C.pango_alignment_get_type()), marshalAlignment},
		{glib.Type(C.pango_ellipsize_mode_get_type()), marshalEllipsizeMode},
		{glib.Type(C.pango_wrap_mode_get_type()), marshalWrapMode},

		// Objects/Interfaces
		//		{glib.Type(C.pango_layout_get_type()), marshalLayout},
	}
	glib_impl.RegisterGValueMarshalers(tm)
}

// Layout is a representation of PangoLayout.
type layout struct {
	pangoLayout *C.PangoLayout
}

// Native returns a pointer to the underlying PangoLayout.
func (v *layout) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *layout) native() *C.PangoLayout {
	return (*C.PangoLayout)(unsafe.Pointer(v.pangoLayout))
}

// LayoutLine is a representation of PangoLayoutLine.
type layoutLine struct {
	pangoLayoutLine *C.PangoLayout
}

// Native returns a pointer to the underlying PangoLayoutLine.
func (v *layoutLine) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *layoutLine) native() *C.PangoLayoutLine {
	return (*C.PangoLayoutLine)(unsafe.Pointer(v.pangoLayoutLine))
}

func init() {
	pango.ALIGN_LEFT = C.PANGO_ALIGN_LEFT
	pango.ALIGN_CENTER = C.PANGO_ALIGN_CENTER
	pango.ALIGN_RIGHT = C.PANGO_ALIGN_RIGHT
	pango.WRAP_WORD = C.PANGO_WRAP_WORD
	pango.WRAP_CHAR = C.PANGO_WRAP_CHAR
	pango.WRAP_WORD_CHAR = C.PANGO_WRAP_WORD_CHAR
	pango.ELLIPSIZE_NONE = C.PANGO_ELLIPSIZE_NONE
	pango.ELLIPSIZE_START = C.PANGO_ELLIPSIZE_START
	pango.ELLIPSIZE_MIDDLE = C.PANGO_ELLIPSIZE_MIDDLE
	pango.ELLIPSIZE_END = C.PANGO_ELLIPSIZE_END
}

func marshalAlignment(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return pango.Alignment(c), nil
}

func marshalWrapMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return pango.WrapMode(c), nil
}

func marshalEllipsizeMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return pango.EllipsizeMode(c), nil
}

/*
func marshalLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

func wrapLayout(obj *glib.Object) *Layout {
	return &Layout{obj}
}
*/

//PangoLayout *pango_layout_new            (PangoContext   *context);
func LayoutNew(context *context) *layout {
	c := C.pango_layout_new(context.native())

	layout := new(layout)
	layout.pangoLayout = (*C.PangoLayout)(c)
	return layout
}

//PangoLayout *pango_layout_copy           (PangoLayout    *src);
func (v *layout) Copy() pango.Layout {
	c := C.pango_layout_copy(v.native())

	layout := new(layout)
	layout.pangoLayout = (*C.PangoLayout)(c)
	return layout
}

//PangoContext  *pango_layout_get_context    (PangoLayout    *layout);
func (v *layout) GetContext() pango.Context {
	c := C.pango_layout_get_context(v.native())

	context := new(context)
	context.pangoContext = (*C.PangoContext)(c)

	return context
}

//void           pango_layout_set_attributes (PangoLayout    *layout,
//					    PangoAttrList  *attrs);
func (v *layout) SetAttributes(attrs pango.AttrList) {
	C.pango_layout_set_attributes(v.native(), toAttrList(attrs).native())
}

//PangoAttrList *pango_layout_get_attributes (PangoLayout    *layout);
func (v *layout) GetAttributes() pango.AttrList {
	c := C.pango_layout_get_attributes(v.native())

	attrList := new(attrList)
	attrList.pangoAttrList = (*C.PangoAttrList)(c)

	return attrList
}

//void           pango_layout_set_text       (PangoLayout    *layout,
//					    const char     *text,
//					    int             length);
func (v *layout) SetText(text string, length int) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_layout_set_text(v.native(), (*C.char)(cstr), (C.int)(length))
}

//const char    *pango_layout_get_text       (PangoLayout    *layout);
func (v *layout) GetText() string {
	c := C.pango_layout_get_text(v.native())
	return C.GoString((*C.char)(c))
}

//gint           pango_layout_get_character_count (PangoLayout *layout);
func (v *layout) GetCharacterCount() int {
	c := C.pango_layout_get_character_count(v.native())
	return int(c)
}

//void           pango_layout_set_markup     (PangoLayout    *layout,
//					    const char     *markup,
//					    int             length);
func (v *layout) SetMarkup(text string, length int) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_layout_set_markup(v.native(), (*C.char)(cstr), (C.int)(length))
}

//void           pango_layout_set_markup_with_accel (PangoLayout    *layout,
//						   const char     *markup,
//						   int             length,
//						   gunichar        accel_marker,
//						   gunichar       *accel_char);

/*
func (v *Layout)SetMarkupWithAccel (text string, length int, accel_marker, accel_char rune){
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_layout_set_markup_with_accel (v.native(),  (*C.char)(cstr), (C.int)(length), (C.gunichar)(accel_marker), (C.gunichar)(accel_char) )
}
*/

//void           pango_layout_set_font_description (PangoLayout                *layout,
//						  const PangoFontDescription *desc);

func (v *layout) SetFontDescription(desc pango.FontDescription) {
	C.pango_layout_set_font_description(v.native(), toFontDescription(desc).native())
}

//const PangoFontDescription *pango_layout_get_font_description (PangoLayout *layout);

func (v *layout) GetFontDescription() pango.FontDescription {
	c := C.pango_layout_get_font_description(v.native())

	desc := new(fontDescription)
	desc.pangoFontDescription = (*C.PangoFontDescription)(c)

	return desc
}

//void           pango_layout_set_width            (PangoLayout                *layout,
//						  int                         width);

func (v *layout) SetWidth(width int) {
	C.pango_layout_set_width(v.native(), C.int(width))
}

//int            pango_layout_get_width            (PangoLayout                *layout);

func (v *layout) GetWidth() int {
	c := C.pango_layout_get_width(v.native())
	return int(c)
}

//void           pango_layout_set_height           (PangoLayout                *layout,
//						  int                         height);

func (v *layout) SetHeight(width int) {
	C.pango_layout_set_height(v.native(), C.int(width))
}

//int            pango_layout_get_height           (PangoLayout                *layout);

func (v *layout) GetHeight() int {
	c := C.pango_layout_get_height(v.native())
	return int(c)
}

//void           pango_layout_set_wrap             (PangoLayout                *layout,
//						  PangoWrapMode               wrap);

func (v *layout) SetWrap(wrap pango.WrapMode) {
	C.pango_layout_set_wrap(v.native(), C.PangoWrapMode(wrap))
}

//PangoWrapMode  pango_layout_get_wrap             (PangoLayout                *layout);

func (v *layout) GetWrap() pango.WrapMode {
	c := C.pango_layout_get_wrap(v.native())
	return pango.WrapMode(c)
}

//gboolean       pango_layout_is_wrapped           (PangoLayout                *layout);

func (v *layout) IsWrapped() bool {
	c := C.pango_layout_is_wrapped(v.native())
	return gobool(c)
}

//void           pango_layout_set_indent           (PangoLayout                *layout,
//						  int                         indent);

func (v *layout) SetIndent(indent int) {
	C.pango_layout_set_indent(v.native(), C.int(indent))
}

//int            pango_layout_get_indent           (PangoLayout                *layout);

func (v *layout) GetIndent() int {
	c := C.pango_layout_get_indent(v.native())
	return int(c)
}

//void           pango_layout_set_spacing          (PangoLayout                *layout,
//						  int                         spacing);
//int            pango_layout_get_spacing          (PangoLayout                *layout);
//void           pango_layout_set_justify          (PangoLayout                *layout,
//						  gboolean                    justify);
//gboolean       pango_layout_get_justify          (PangoLayout                *layout);
//void           pango_layout_set_auto_dir         (PangoLayout                *layout,
//						  gboolean                    auto_dir);
//gboolean       pango_layout_get_auto_dir         (PangoLayout                *layout);
//void           pango_layout_set_alignment        (PangoLayout                *layout,
//						  PangoAlignment              alignment);
//PangoAlignment pango_layout_get_alignment        (PangoLayout                *layout);
//
//void           pango_layout_set_tabs             (PangoLayout                *layout,
//						  PangoTabArray              *tabs);
//
//PangoTabArray* pango_layout_get_tabs             (PangoLayout                *layout);
//
//void           pango_layout_set_single_paragraph_mode (PangoLayout                *layout,
//						       gboolean                    setting);
//gboolean       pango_layout_get_single_paragraph_mode (PangoLayout                *layout);
//
//void               pango_layout_set_ellipsize (PangoLayout        *layout,
//					       PangoEllipsizeMode  ellipsize);
//PangoEllipsizeMode pango_layout_get_ellipsize (PangoLayout        *layout);
//gboolean           pango_layout_is_ellipsized (PangoLayout        *layout);
//
//int      pango_layout_get_unknown_glyphs_count (PangoLayout    *layout);
//
//void     pango_layout_context_changed (PangoLayout    *layout);
//guint    pango_layout_get_serial      (PangoLayout    *layout);
//
//void     pango_layout_get_log_attrs (PangoLayout    *layout,
//				     PangoLogAttr  **attrs,
//				     gint           *n_attrs);
//
//const PangoLogAttr *pango_layout_get_log_attrs_readonly (PangoLayout *layout,
//							 gint        *n_attrs);
//
//void     pango_layout_index_to_pos         (PangoLayout    *layout,
//					    int             index_,
//					    PangoRectangle *pos);
//void     pango_layout_index_to_line_x      (PangoLayout    *layout,
//					    int             index_,
//					    gboolean        trailing,
//					    int            *line,
//					    int            *x_pos);
//void     pango_layout_get_cursor_pos       (PangoLayout    *layout,
//					    int             index_,
//					    PangoRectangle *strong_pos,
//					    PangoRectangle *weak_pos);
//void     pango_layout_move_cursor_visually (PangoLayout    *layout,
//					    gboolean        strong,
//					    int             old_index,
//					    int             old_trailing,
//					    int             direction,
//					    int            *new_index,
//					    int            *new_trailing);
//gboolean pango_layout_xy_to_index          (PangoLayout    *layout,
//					    int             x,
//					    int             y,
//					    int            *index_,
//					    int            *trailing);
//void     pango_layout_get_extents          (PangoLayout    *layout,
//					    PangoRectangle *ink_rect,
//					    PangoRectangle *logical_rect);
//void     pango_layout_get_pixel_extents    (PangoLayout    *layout,
//					    PangoRectangle *ink_rect,
//					    PangoRectangle *logical_rect);

//void     pango_layout_get_size             (PangoLayout    *layout,
//					    int            *width,
//					    int            *height);
func (v *layout) GetSize() (int, int) {
	var w, h C.int
	C.pango_layout_get_size(v.native(), &w, &h)
	return int(w), int(h)
}

//void     pango_layout_get_pixel_size       (PangoLayout    *layout,
//					    int            *width,
//					    int            *height);
//int      pango_layout_get_baseline         (PangoLayout    *layout);
//
//int              pango_layout_get_line_count       (PangoLayout    *layout);
//PangoLayoutLine *pango_layout_get_line             (PangoLayout    *layout,
//						    int             line);
//PangoLayoutLine *pango_layout_get_line_readonly    (PangoLayout    *layout,
//						    int             line);
//GSList *         pango_layout_get_lines            (PangoLayout    *layout);
//GSList *         pango_layout_get_lines_readonly   (PangoLayout    *layout);
//
//
//#define PANGO_TYPE_LAYOUT_LINE (pango_layout_line_get_type ())
//
//GType    pango_layout_line_get_type     (void) G_GNUC_CONST;
//
//PangoLayoutLine *pango_layout_line_ref   (PangoLayoutLine *line);
//void             pango_layout_line_unref (PangoLayoutLine *line);
//
//gboolean pango_layout_line_x_to_index   (PangoLayoutLine  *line,
//					 int               x_pos,
//					 int              *index_,
//					 int              *trailing);
//void     pango_layout_line_index_to_x   (PangoLayoutLine  *line,
//					 int               index_,
//					 gboolean          trailing,
//					 int              *x_pos);
//void     pango_layout_line_get_x_ranges (PangoLayoutLine  *line,
//					 int               start_index,
//					 int               end_index,
//					 int             **ranges,
//					 int              *n_ranges);
//void     pango_layout_line_get_extents  (PangoLayoutLine  *line,
//					 PangoRectangle   *ink_rect,
//					 PangoRectangle   *logical_rect);
//void     pango_layout_line_get_pixel_extents (PangoLayoutLine *layout_line,
//					      PangoRectangle  *ink_rect,
//					      PangoRectangle  *logical_rect);
//
//typedef struct _PangoLayoutIter PangoLayoutIter;
//
//#define PANGO_TYPE_LAYOUT_ITER         (pango_layout_iter_get_type ())
//
//GType            pango_layout_iter_get_type (void) G_GNUC_CONST;
//
//PangoLayoutIter *pango_layout_get_iter  (PangoLayout     *layout);
//PangoLayoutIter *pango_layout_iter_copy (PangoLayoutIter *iter);
//void             pango_layout_iter_free (PangoLayoutIter *iter);
//
//int              pango_layout_iter_get_index  (PangoLayoutIter *iter);
//PangoLayoutRun  *pango_layout_iter_get_run    (PangoLayoutIter *iter);
//PangoLayoutRun  *pango_layout_iter_get_run_readonly   (PangoLayoutIter *iter);
//PangoLayoutLine *pango_layout_iter_get_line   (PangoLayoutIter *iter);
//PangoLayoutLine *pango_layout_iter_get_line_readonly  (PangoLayoutIter *iter);
//gboolean         pango_layout_iter_at_last_line (PangoLayoutIter *iter);
//PangoLayout     *pango_layout_iter_get_layout (PangoLayoutIter *iter);
//
//gboolean pango_layout_iter_next_char    (PangoLayoutIter *iter);
//gboolean pango_layout_iter_next_cluster (PangoLayoutIter *iter);
//gboolean pango_layout_iter_next_run     (PangoLayoutIter *iter);
//gboolean pango_layout_iter_next_line    (PangoLayoutIter *iter);
//
//void pango_layout_iter_get_char_extents    (PangoLayoutIter *iter,
//					    PangoRectangle  *logical_rect);
//void pango_layout_iter_get_cluster_extents (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
//void pango_layout_iter_get_run_extents     (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
//void pango_layout_iter_get_line_extents    (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
/* All the yranges meet, unlike the logical_rect's (i.e. the yranges
 * assign between-line spacing to the nearest line)
 */
//void pango_layout_iter_get_line_yrange     (PangoLayoutIter *iter,
//					    int             *y0_,
//					    int             *y1_);
//void pango_layout_iter_get_layout_extents  (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
//int  pango_layout_iter_get_baseline        (PangoLayoutIter *iter);
//