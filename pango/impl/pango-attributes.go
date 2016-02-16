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
		{glib.Type(C.pango_attr_type_get_type()), marshalAttrType},
		{glib.Type(C.pango_underline_get_type()), marshalUnderline},
	}
	glib_impl.RegisterGValueMarshalers(tm)
}

/* PangoColor */

// Color is a representation of PangoColor.
type color struct {
	pangoColor *C.PangoColor
}

// Native returns a pointer to the underlying PangoColor.
func (v *color) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *color) native() *C.PangoColor {
	return (*C.PangoColor)(unsafe.Pointer(v.pangoColor))
}

func (v *color) Set(red, green, blue uint16) {
	v.native().red = C.guint16(red)
	v.native().green = C.guint16(green)
	v.native().blue = C.guint16(blue)
}

func (v *color) Get() (red, green, blue uint16) {
	return uint16(v.native().red), uint16(v.native().green), uint16(v.native().blue)
}

//PangoColor *pango_color_copy     (const PangoColor *src);
func (v *color) Copy(c pango.Color) pango.Color {
	w := new(color)
	w.pangoColor = C.pango_color_copy(v.native())
	return w
}

//void        pango_color_free     (PangoColor       *color);
func (v *color) Free() {
	C.pango_color_free(v.native())
}

//gboolean    pango_color_parse    (PangoColor       *color,
//			  const char       *spec);
func (v *color) Parse(spec string) bool {
	cstr := C.CString(spec)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_color_parse(v.native(), (*C.char)(cstr))
	return gobool(c)
}

//gchar      *pango_color_to_string(const PangoColor *color);
func (v *color) ToString() string {
	c := C.pango_color_to_string(v.native())
	return C.GoString((*C.char)(c))
}

/* ---  ---  --- Attributes ---  ---  ---  */

// AttrList is a representation of PangoAttrList.
type attrList struct {
	pangoAttrList *C.PangoAttrList
}

// Native returns a pointer to the underlying PangoLayout.
func (v attrList) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v attrList) native() *C.PangoAttrList {
	return (*C.PangoAttrList)(unsafe.Pointer(v.pangoAttrList))
}

func init() {
	pango.ATTR_INVALID = C.PANGO_ATTR_INVALID                         /* 0 is an invalid attribute type */
	pango.ATTR_LANGUAGE = C.PANGO_ATTR_LANGUAGE                       /* PangoAttrLanguage */
	pango.ATTR_FAMILY = C.PANGO_ATTR_FAMILY                           /* PangoAttrString */
	pango.ATTR_STYLE = C.PANGO_ATTR_STYLE                             /* PangoAttrInt */
	pango.ATTR_WEIGHT = C.PANGO_ATTR_WEIGHT                           /* PangoAttrInt */
	pango.ATTR_VARIANT = C.PANGO_ATTR_VARIANT                         /* PangoAttrInt */
	pango.ATTR_STRETCH = C.PANGO_ATTR_STRETCH                         /* PangoAttrInt */
	pango.ATTR_SIZE = C.PANGO_ATTR_SIZE                               /* PangoAttrSize */
	pango.ATTR_FONT_DESC = C.PANGO_ATTR_FONT_DESC                     /* PangoAttrFontDesc */
	pango.ATTR_FOREGROUND = C.PANGO_ATTR_FOREGROUND                   /* PangoAttrColor */
	pango.ATTR_BACKGROUND = C.PANGO_ATTR_BACKGROUND                   /* PangoAttrColor */
	pango.ATTR_UNDERLINE = C.PANGO_ATTR_UNDERLINE                     /* PangoAttrInt */
	pango.ATTR_STRIKETHROUGH = C.PANGO_ATTR_STRIKETHROUGH             /* PangoAttrInt */
	pango.ATTR_RISE = C.PANGO_ATTR_RISE                               /* PangoAttrInt */
	pango.ATTR_SHAPE = C.PANGO_ATTR_SHAPE                             /* PangoAttrShape */
	pango.ATTR_SCALE = C.PANGO_ATTR_SCALE                             /* PangoAttrFloat */
	pango.ATTR_FALLBACK = C.PANGO_ATTR_FALLBACK                       /* PangoAttrInt */
	pango.ATTR_LETTER_SPACING = C.PANGO_ATTR_LETTER_SPACING           /* PangoAttrInt */
	pango.ATTR_UNDERLINE_COLOR = C.PANGO_ATTR_UNDERLINE_COLOR         /* PangoAttrColor */
	pango.ATTR_STRIKETHROUGH_COLOR = C.PANGO_ATTR_STRIKETHROUGH_COLOR /* PangoAttrColor */
	pango.ATTR_ABSOLUTE_SIZE = C.PANGO_ATTR_ABSOLUTE_SIZE             /* PangoAttrSize */
	pango.ATTR_GRAVITY = C.PANGO_ATTR_GRAVITY                         /* PangoAttrInt */
	pango.ATTR_GRAVITY_HINT = C.PANGO_ATTR_GRAVITY_HINT               /* PangoAttrInt */

	pango.UNDERLINE_NONE = C.PANGO_UNDERLINE_NONE
	pango.UNDERLINE_SINGLE = C.PANGO_UNDERLINE_SINGLE
	pango.UNDERLINE_DOUBLE = C.PANGO_UNDERLINE_DOUBLE
	pango.UNDERLINE_LOW = C.PANGO_UNDERLINE_LOW
	pango.UNDERLINE_ERROR = C.PANGO_UNDERLINE_ERROR

	pango.ATTR_INDEX_FROM_TEXT_BEGINNING = 0
	pango.ATTR_INDEX_TO_TEXT_END = C.G_MAXUINT
}

func marshalAttrType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return pango.AttrType(c), nil
}

func marshalUnderline(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return pango.Underline(c), nil
}

// Attribute is a representation of Pango's PangoAttribute.
type attribute struct {
	pangoAttribute *C.PangoAttribute
	//start_index, end_index uint
}

// Native returns a pointer to the underlying PangoColor.
func (v *attribute) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *attribute) native() *C.PangoAttribute {
	return (*C.PangoAttribute)(unsafe.Pointer(v.pangoAttribute))
}

/*
//typedef gboolean (*PangoAttrFilterFunc) (PangoAttribute *attribute,
//					 gpointer        user_data);
func (v *Attribute) AttrFilterFunc(user_data uintptr) bool {
	c := C.PangoAttrFilterFunc(Attribute.native(), C.gpointer(user_data))
	return gobool(c)
}

//typedef gpointer (*PangoAttrDataCopyFunc) (gconstpointer user_data);
func AttrDataCopyFunc(user_data uintptr) uintptr {
	c := C.PangoAttrDataCopyFunc(C.gpointer(user_data))
	return uintptr(c)
}
*/

// AttrClass is a representation of Pango's PangoAttrClass.
type attrClass struct {
	//PangoAttrType type;
}

// AttrString is a representation of Pango's PangoAttrString.
type attrString struct {
	attribute
	//char *value;
}

// AttrLanguage is a representation of Pango's PangoAttrLanguage.
type attrLanguage struct {
	attribute
	//PangoLanguage *value;
}

// AttrInt is a representation of Pango's PangoAttrInt.
type attrInt struct {
	attribute
	//int value;
}

// AttrFloat is a representation of Pango's PangoAttrFloat.
type attrFloat struct {
	attribute
	//double value;
}

// AttrColor is a representation of Pango's AttrColor.
type attrColor struct {
	attribute
	color
}

// AttrSize is a representation of Pango's PangoAttrSize.
type attrSize struct {
	attribute
	//int size;
	//guint absolute : 1;
}

// AttrShape is a representation of Pango's PangoAttrShape.
type attrShape struct {
	attribute
	//PangoRectangle ink_rect;
	//PangoRectangle logical_rect;

	//gpointer              data;
	//PangoAttrDataCopyFunc copy_func;
	//GDestroyNotify        destroy_func;
}

// AttrFontDesc is a representation of Pango's PangoAttrFontDesc.
type attrFontDesc struct {
	attribute
	//PangoFontDescription *desc;
}

/*
PangoAttrType         pango_attr_type_register (const gchar        *name);
const char *          pango_attr_type_get_name (PangoAttrType       type) G_GNUC_CONST;

void             pango_attribute_init        (PangoAttribute       *attr,
					      const PangoAttrClass *klass);
PangoAttribute * pango_attribute_copy        (const PangoAttribute *attr);
void             pango_attribute_destroy     (PangoAttribute       *attr);
gboolean         pango_attribute_equal       (const PangoAttribute *attr1,
					      const PangoAttribute *attr2) G_GNUC_PURE;

PangoAttribute *pango_attr_language_new      (PangoLanguage              *language);
PangoAttribute *pango_attr_family_new        (const char                 *family);
PangoAttribute *pango_attr_foreground_new    (guint16                     red,
					      guint16                     green,
					      guint16                     blue);
PangoAttribute *pango_attr_background_new    (guint16                     red,
					      guint16                     green,
					      guint16                     blue);
PangoAttribute *pango_attr_size_new          (int                         size);
PangoAttribute *pango_attr_size_new_absolute (int                         size);
PangoAttribute *pango_attr_style_new         (PangoStyle                  style);
PangoAttribute *pango_attr_weight_new        (PangoWeight                 weight);
PangoAttribute *pango_attr_variant_new       (PangoVariant                variant);
PangoAttribute *pango_attr_stretch_new       (PangoStretch                stretch);
PangoAttribute *pango_attr_font_desc_new     (const PangoFontDescription *desc);

PangoAttribute *pango_attr_underline_new           (PangoUnderline underline);
PangoAttribute *pango_attr_underline_color_new     (guint16        red,
						    guint16        green,
						    guint16        blue);
PangoAttribute *pango_attr_strikethrough_new       (gboolean       strikethrough);
PangoAttribute *pango_attr_strikethrough_color_new (guint16        red,
						    guint16        green,
						    guint16        blue);

PangoAttribute *pango_attr_rise_new          (int                         rise);
PangoAttribute *pango_attr_scale_new         (double                      scale_factor);
PangoAttribute *pango_attr_fallback_new      (gboolean                    enable_fallback);
PangoAttribute *pango_attr_letter_spacing_new (int                        letter_spacing);

PangoAttribute *pango_attr_shape_new           (const PangoRectangle       *ink_rect,
						const PangoRectangle       *logical_rect);
PangoAttribute *pango_attr_shape_new_with_data (const PangoRectangle       *ink_rect,
						const PangoRectangle       *logical_rect,
						gpointer                    data,
						PangoAttrDataCopyFunc       copy_func,
						GDestroyNotify              destroy_func);

PangoAttribute *pango_attr_gravity_new      (PangoGravity     gravity);
PangoAttribute *pango_attr_gravity_hint_new (PangoGravityHint hint);

GType              pango_attr_list_get_type      (void) G_GNUC_CONST;
PangoAttrList *    pango_attr_list_new           (void);
PangoAttrList *    pango_attr_list_ref           (PangoAttrList  *list);
void               pango_attr_list_unref         (PangoAttrList  *list);
PangoAttrList *    pango_attr_list_copy          (PangoAttrList  *list);
void               pango_attr_list_insert        (PangoAttrList  *list,
						  PangoAttribute *attr);
void               pango_attr_list_insert_before (PangoAttrList  *list,
						  PangoAttribute *attr);
void               pango_attr_list_change        (PangoAttrList  *list,
						  PangoAttribute *attr);
void               pango_attr_list_splice        (PangoAttrList  *list,
						  PangoAttrList  *other,
						  gint            pos,
						  gint            len);

PangoAttrList *pango_attr_list_filter (PangoAttrList       *list,
				       PangoAttrFilterFunc  func,
				       gpointer             data);

PangoAttrIterator *pango_attr_list_get_iterator  (PangoAttrList  *list);

void               pango_attr_iterator_range    (PangoAttrIterator     *iterator,
						 gint                  *start,
						 gint                  *end);
gboolean           pango_attr_iterator_next     (PangoAttrIterator     *iterator);
PangoAttrIterator *pango_attr_iterator_copy     (PangoAttrIterator     *iterator);
void               pango_attr_iterator_destroy  (PangoAttrIterator     *iterator);
PangoAttribute *   pango_attr_iterator_get      (PangoAttrIterator     *iterator,
						 PangoAttrType          type);
void               pango_attr_iterator_get_font (PangoAttrIterator     *iterator,
						 PangoFontDescription  *desc,
						 PangoLanguage        **language,
						 GSList               **extra_attrs);
GSList *          pango_attr_iterator_get_attrs (PangoAttrIterator     *iterator);


gboolean pango_parse_markup (const char                 *markup_text,
			     int                         length,
			     gunichar                    accel_marker,
			     PangoAttrList             **attr_list,
			     char                      **text,
			     gunichar                   *accel_char,
			     GError                    **error);

GMarkupParseContext * pango_markup_parser_new (gunichar               accel_marker);
gboolean              pango_markup_parser_finish (GMarkupParseContext   *context,
                                                  PangoAttrList        **attr_list,
                                                  char                 **text,
                                                  gunichar              *accel_char,
                                                  GError               **error);
*/
