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
		// Objects/Interfaces
		{glib.Type(C.pango_font_description_get_type()), marshalFontDescription},
	}
	glib_impl.RegisterGValueMarshalers(tm)
}

// FontDescription is a representation of PangoFontDescription.
type FontDescription struct {
	pangoFontDescription *C.PangoFontDescription
}

// Native returns a pointer to the underlying PangoLayout.
func (v *FontDescription) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *FontDescription) native() *C.PangoFontDescription {
	return (*C.PangoFontDescription)(unsafe.Pointer(v.pangoFontDescription))
}

// FontMetrics is a representation of PangoFontMetrics.
type FontMetrics struct {
	pangoFontMetrics *C.PangoFontMetrics
}

// Native returns a pointer to the underlying PangoLayout.
func (v *FontMetrics) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *FontMetrics) native() *C.PangoFontMetrics {
	return (*C.PangoFontMetrics)(unsafe.Pointer(v.pangoFontMetrics))
}

func init() {
	pango.PANGO_SCALE = C.PANGO_SCALE

	pango.STYLE_NORMAL = C.PANGO_STYLE_NORMAL
	pango.STYLE_OBLIQUE = C.PANGO_STYLE_OBLIQUE
	pango.STYLE_ITALIC = C.PANGO_STYLE_ITALIC

	pango.VARIANT_NORMAL = C.PANGO_VARIANT_NORMAL
	pango.VARIANT_SMALL_CAPS = C.PANGO_VARIANT_SMALL_CAPS

	pango.WEIGHT_THIN = C.PANGO_WEIGHT_THIN             /* 100 */
	pango.WEIGHT_ULTRALIGHT = C.PANGO_WEIGHT_ULTRALIGHT /* 200 */
	pango.WEIGHT_LIGHT = C.PANGO_WEIGHT_LIGHT           /* 300 */
	pango.WEIGHT_SEMILIGHT = 350                        /* 350 */
	pango.WEIGHT_BOOK = C.PANGO_WEIGHT_BOOK             /* 380 */
	pango.WEIGHT_NORMAL = C.PANGO_WEIGHT_NORMAL         /* 400 */
	pango.WEIGHT_MEDIUM = C.PANGO_WEIGHT_MEDIUM         /* 500 */
	pango.WEIGHT_SEMIBOLD = C.PANGO_WEIGHT_SEMIBOLD     /* 600 */
	pango.WEIGHT_BOLD = C.PANGO_WEIGHT_BOLD             /* 700 */
	pango.WEIGHT_ULTRABOLD = C.PANGO_WEIGHT_ULTRABOLD   /* 800 */
	pango.WEIGHT_HEAVY = C.PANGO_WEIGHT_HEAVY           /* 900 */
	pango.WEIGHT_ULTRAHEAVY = C.PANGO_WEIGHT_ULTRAHEAVY /* 1000 */

	pango.STRETCH_ULTRA_CONDENSED = C.PANGO_STRETCH_ULTRA_CONDENSED
	pango.STRETCH_EXTRA_CONDENSEDStretch = C.PANGO_STRETCH_EXTRA_CONDENSED
	pango.STRETCH_CONDENSEDStretch = C.PANGO_STRETCH_CONDENSED
	pango.STRETCH_SEMI_CONDENSEDStretch = C.PANGO_STRETCH_SEMI_CONDENSED
	pango.STRETCH_NORMALStretch = C.PANGO_STRETCH_NORMAL
	pango.STRETCH_SEMI_EXPANDEDStretch = C.PANGO_STRETCH_SEMI_EXPANDED
	pango.STRETCH_EXPANDEDStretch = C.PANGO_STRETCH_EXPANDED
	pango.STRETCH_EXTRA_EXPANDEDStretch = C.PANGO_STRETCH_EXTRA_EXPANDED
	pango.STRETCH_ULTRA_EXPANDEDStretch = C.PANGO_STRETCH_ULTRA_EXPANDED

	pango.FONT_MASK_FAMILY = C.PANGO_FONT_MASK_FAMILY           /*  1 << 0 */
	pango.FONT_MASK_STYLEFontMask = C.PANGO_FONT_MASK_STYLE     /*  1 << 1 */
	pango.FONT_MASK_VARIANTFontMask = C.PANGO_FONT_MASK_VARIANT /*  1 << 2 */
	pango.FONT_MASK_WEIGHTFontMask = C.PANGO_FONT_MASK_WEIGHT   /*  1 << 3 */
	pango.FONT_MASK_STRETCHFontMask = C.PANGO_FONT_MASK_STRETCH /*  1 << 4 */
	pango.FONT_MASK_SIZEFontMask = C.PANGO_FONT_MASK_SIZE       /*  1 << 5 */
	pango.FONT_MASK_GRAVITYFontMask = C.PANGO_FONT_MASK_GRAVITY /*  1 << 6 */

	pango.SCALE_XX_SMALL = /* C.PANGO_SCALE_XX_SMALL */ 0.5787037037037
	pango.SCALE_X_SMALL = /*C.PANGO_SCALE_X_SMALL  */ 0.6444444444444
	pango.SCALE_SMALL = /*C.PANGO_SCALE_SMALL    */ 0.8333333333333
	pango.SCALE_MEDIUM = /*C.PANGO_SCALE_MEDIUM   */ 1.0
	pango.SCALE_LARGE = /*C.PANGO_SCALE_LARGE    */ 1.2
	pango.SCALE_X_LARGE = /*C.PANGO_SCALE_X_LARGE  */ 1.4399999999999
	pango.SCALE_XX_LARGE = /*C.PANGO_SCALE_XX_LARGE */ 1.728
}

/*
 * PangoFontDescription
 */

func marshalFontDescription(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c2 := (*C.PangoFontDescription)(unsafe.Pointer(c))
	return wrapFontDescription(c2), nil
}

func wrapFontDescription(obj *C.PangoFontDescription) *FontDescription {
	return &FontDescription{obj}
}

//PangoFontDescription *pango_font_description_new         (void);
func FontDescriptionNew() *FontDescription {
	c := C.pango_font_description_new()
	v := new(FontDescription)
	v.pangoFontDescription = c
	return v
}

//PangoFontDescription *pango_font_description_copy        (const PangoFontDescription  *desc);
func (v *FontDescription) Copy() pango.FontDescription {
	c := C.pango_font_description_copy(v.native())
	v2 := new(FontDescription)
	v2.pangoFontDescription = c
	return v2
}

//PangoFontDescription *pango_font_description_copy_static (const PangoFontDescription  *desc);
func (v *FontDescription) CopyStatic() pango.FontDescription {
	c := C.pango_font_description_copy_static(v.native())
	v2 := new(FontDescription)
	v2.pangoFontDescription = c
	return v2
}

//guint                 pango_font_description_hash        (const PangoFontDescription  *desc) G_GNUC_PURE;
func (v *FontDescription) Hash() uint {
	c := C.pango_font_description_hash(v.native())
	return uint(c)
}

//gboolean              pango_font_description_equal       (const PangoFontDescription  *desc1,
//							  const PangoFontDescription  *desc2) G_GNUC_PURE;
func (v *FontDescription) Equal(v2 pango.FontDescription) bool {
	c := C.pango_font_description_equal(v.native(), v2.(*FontDescription).native())
	return gobool(c)
}

//void                  pango_font_description_free        (PangoFontDescription        *desc);
func (v *FontDescription) Free() {
	C.pango_font_description_free(v.native())
}

//void                  pango_font_descriptions_free       (PangoFontDescription       **descs,
//							  int                          n_descs);
//func (v *FontDescription) FontDescriptionsFree(n_descs int) {
//	C.pango_font_descriptions_free(v.native(), C.int(n_descs))
//}

//void                 pango_font_description_set_family        (PangoFontDescription *desc,
//							       const char           *family);
func (v *FontDescription) SetFamily(family string) {
	cstr := C.CString(family)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_font_description_set_family(v.native(), (*C.char)(cstr))
}

//void                 pango_font_description_set_family_static (PangoFontDescription *desc,
//							       const char           *family);
func (v *FontDescription) SetFamilyStatic(family string) {
	cstr := C.CString(family)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_font_description_set_family_static(v.native(), (*C.char)(cstr))
}

//const char          *pango_font_description_get_family        (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetFamily() string {
	c := C.pango_font_description_get_family(v.native())
	return C.GoString((*C.char)(c))
}

//void                 pango_font_description_set_style         (PangoFontDescription *desc,
//							       PangoStyle            style);
func (v *FontDescription) SetStyle(style pango.Style) {
	C.pango_font_description_set_style(v.native(), (C.PangoStyle)(style))
}

//PangoStyle           pango_font_description_get_style         (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetStyle() pango.Style {
	c := C.pango_font_description_get_style(v.native())
	return pango.Style(c)
}

//void                 pango_font_description_set_variant       (PangoFontDescription *desc,
//							       PangoVariant          variant);
//PangoVariant         pango_font_description_get_variant       (const PangoFontDescription *desc) G_GNUC_PURE;

//void                 pango_font_description_set_weight        (PangoFontDescription *desc,
//							       PangoWeight           weight);
func (v *FontDescription) SetWeight(weight pango.Weight) {
	C.pango_font_description_set_weight(v.native(), (C.PangoWeight)(weight))
}

//PangoWeight          pango_font_description_get_weight        (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetWeight() pango.Weight {
	c := C.pango_font_description_get_weight(v.native())
	return pango.Weight(c)
}

//void                 pango_font_description_set_stretch       (PangoFontDescription *desc,
//							       PangoStretch          stretch);
func (v *FontDescription) SetStretch(stretch pango.Stretch) {
	C.pango_font_description_set_stretch(v.native(), (C.PangoStretch)(stretch))
}

//PangoStretch         pango_font_description_get_stretch       (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetStretch() pango.Stretch {
	c := C.pango_font_description_get_stretch(v.native())
	return pango.Stretch(c)
}

//void                 pango_font_description_set_size          (PangoFontDescription *desc,
//							       gint                  size);
func (v *FontDescription) SetSize(size int) {
	C.pango_font_description_set_size(v.native(), (C.gint)(size))
}

//gint                 pango_font_description_get_size          (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetSize() int {
	c := C.pango_font_description_get_size(v.native())
	return int(c)
}

//void                 pango_font_description_set_absolute_size (PangoFontDescription *desc,
//							       double                size);
func (v *FontDescription) SetAbsoluteSize(size float64) {
	C.pango_font_description_set_absolute_size(v.native(), (C.double)(size))
}

//gboolean             pango_font_description_get_size_is_absolute (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetSizeIsAbsolute() bool {
	c := C.pango_font_description_get_size_is_absolute(v.native())
	return gobool(c)
}

//void                 pango_font_description_set_gravity       (PangoFontDescription *desc,
//							       PangoGravity          gravity);
func (v *FontDescription) SetGravity(gravity pango.Gravity) {
	C.pango_font_description_set_gravity(v.native(), (C.PangoGravity)(gravity))
}

//PangoGravity         pango_font_description_get_gravity       (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetGravity() pango.Gravity {
	c := C.pango_font_description_get_gravity(v.native())
	return pango.Gravity(c)
}

//PangoFontMask pango_font_description_get_set_fields (const PangoFontDescription *desc) G_GNUC_PURE;
func (v *FontDescription) GetSetFields() pango.FontMask {
	c := C.pango_font_description_get_set_fields(v.native())
	return pango.FontMask(c)
}

//void          pango_font_description_unset_fields   (PangoFontDescription       *desc,
//						     PangoFontMask               to_unset);
func (v *FontDescription) GetUnsetFields(to_unset pango.FontMask) {
	C.pango_font_description_unset_fields(v.native(), (C.PangoFontMask)(to_unset))
}

//void pango_font_description_merge        (PangoFontDescription       *desc,
//					  const PangoFontDescription *desc_to_merge,
//					  gboolean                    replace_existing);
func (v *FontDescription) Merge(desc_to_merge pango.FontDescription, replace_existing bool) {
	C.pango_font_description_merge(v.native(), desc_to_merge.(*FontDescription).native(), gbool(replace_existing))
}

//void pango_font_description_merge_static (PangoFontDescription       *desc,
//					  const PangoFontDescription *desc_to_merge,
//					  gboolean                    replace_existing);
func (v *FontDescription) MergeStatic(desc_to_merge pango.FontDescription, replace_existing bool) {
	C.pango_font_description_merge_static(v.native(), desc_to_merge.(*FontDescription).native(), gbool(replace_existing))
}

//gboolean pango_font_description_better_match (const PangoFontDescription *desc,
//					      const PangoFontDescription *old_match,
//					      const PangoFontDescription *new_match) G_GNUC_PURE;
func (v *FontDescription) BetterMatch(old_match, new_match pango.FontDescription) bool {
	c := C.pango_font_description_better_match(v.native(), old_match.(*FontDescription).native(), new_match.(*FontDescription).native())
	return gobool(c)
}

//PangoFontDescription *pango_font_description_from_string (const char                  *str);
func FontDescriptionFromString(str string) *FontDescription {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_font_description_from_string((*C.char)(cstr))
	v := new(FontDescription)
	v.pangoFontDescription = c
	return v
}

//char *                pango_font_description_to_string   (const PangoFontDescription  *desc);
func (v *FontDescription) ToString() string {
	c := C.pango_font_description_to_string(v.native())
	return C.GoString((*C.char)(c))
}

//char *                pango_font_description_to_filename (const PangoFontDescription  *desc);
func (v *FontDescription) ToFilename() string {
	c := C.pango_font_description_to_filename(v.native())
	return C.GoString((*C.char)(c))
}

///*
// * PangoFontMetrics
// */
//
///**
// * PANGO_TYPE_FONT_METRICS:
// *
// * The #GObject type for #PangoFontMetrics.
// */
//#define PANGO_TYPE_FONT_METRICS  (pango_font_metrics_get_type ())
//GType             pango_font_metrics_get_type                    (void) G_GNUC_CONST;
//PangoFontMetrics *pango_font_metrics_ref                         (PangoFontMetrics *metrics);
//void              pango_font_metrics_unref                       (PangoFontMetrics *metrics);
//int               pango_font_metrics_get_ascent                  (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_descent                 (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_approximate_char_width  (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_approximate_digit_width (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_underline_position      (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_underline_thickness     (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_strikethrough_position  (PangoFontMetrics *metrics) G_GNUC_PURE;
//int               pango_font_metrics_get_strikethrough_thickness (PangoFontMetrics *metrics) G_GNUC_PURE;
//
//#ifdef PANGO_ENABLE_BACKEND
//
//PangoFontMetrics *pango_font_metrics_new (void);
//
//struct _PangoFontMetrics
//{
//  guint ref_count;
//
//  int ascent;
//  int descent;
//  int approximate_char_width;
//  int approximate_digit_width;
//  int underline_position;
//  int underline_thickness;
//  int strikethrough_position;
//  int strikethrough_thickness;
//};
//
//#endif /* PANGO_ENABLE_BACKEND */
//
///*
// * PangoFontFamily
// */
//
///**
// * PANGO_TYPE_FONT_FAMILY:
// *
// * The #GObject type for #PangoFontFamily.
// */
///**
// * PANGO_FONT_FAMILY:
// * @object: a #GObject.
// *
// * Casts a #GObject to a #PangoFontFamily.
// */
///**
// * PANGO_IS_FONT_FAMILY:
// * @object: a #GObject.
// *
// * Returns: %TRUE if @object is a #PangoFontFamily.
// */
//#define PANGO_TYPE_FONT_FAMILY              (pango_font_family_get_type ())
//#define PANGO_FONT_FAMILY(object)           (G_TYPE_CHECK_INSTANCE_CAST ((object), PANGO_TYPE_FONT_FAMILY, PangoFontFamily))
//#define PANGO_IS_FONT_FAMILY(object)        (G_TYPE_CHECK_INSTANCE_TYPE ((object), PANGO_TYPE_FONT_FAMILY))
//
//typedef struct _PangoFontFamily      PangoFontFamily;
//typedef struct _PangoFontFace        PangoFontFace;
//
//GType      pango_font_family_get_type       (void) G_GNUC_CONST;
//
//void                 pango_font_family_list_faces (PangoFontFamily  *family,
//						   PangoFontFace  ***faces,
//						   int              *n_faces);
//const char *pango_font_family_get_name   (PangoFontFamily  *family) G_GNUC_PURE;
//gboolean   pango_font_family_is_monospace         (PangoFontFamily  *family) G_GNUC_PURE;
//
//#ifdef PANGO_ENABLE_BACKEND
//
//#define PANGO_FONT_FAMILY_CLASS(klass)      (G_TYPE_CHECK_CLASS_CAST ((klass), PANGO_TYPE_FONT_FAMILY, PangoFontFamilyClass))
//#define PANGO_IS_FONT_FAMILY_CLASS(klass)   (G_TYPE_CHECK_CLASS_TYPE ((klass), PANGO_TYPE_FONT_FAMILY))
//#define PANGO_FONT_FAMILY_GET_CLASS(obj)    (G_TYPE_INSTANCE_GET_CLASS ((obj), PANGO_TYPE_FONT_FAMILY, PangoFontFamilyClass))
//
//typedef struct _PangoFontFamilyClass PangoFontFamilyClass;
//
//
///**
// * PangoFontFamily:
// *
// * The #PangoFontFamily structure is used to represent a family of related
// * font faces. The faces in a family share a common design, but differ in
// * slant, weight, width and other aspects.
// */
//struct _PangoFontFamily
//{
//  GObject parent_instance;
//};
//
//struct _PangoFontFamilyClass
//{
//  GObjectClass parent_class;
//
//  /*< public >*/
//
//  void  (*list_faces)      (PangoFontFamily  *family,
//			    PangoFontFace  ***faces,
//			    int              *n_faces);
//  const char * (*get_name) (PangoFontFamily  *family);
//  gboolean (*is_monospace) (PangoFontFamily *family);
//
//  /*< private >*/
//
//  /* Padding for future expansion */
//  void (*_pango_reserved2) (void);
//  void (*_pango_reserved3) (void);
//  void (*_pango_reserved4) (void);
//};
//
//#endif /* PANGO_ENABLE_BACKEND */
//
///*
// * PangoFontFace
// */
//
///**
// * PANGO_TYPE_FONT_FACE:
// *
// * The #GObject type for #PangoFontFace.
// */
///**
// * PANGO_FONT_FACE:
// * @object: a #GObject.
// *
// * Casts a #GObject to a #PangoFontFace.
// */
///**
// * PANGO_IS_FONT_FACE:
// * @object: a #GObject.
// *
// * Returns: %TRUE if @object is a #PangoFontFace.
// */
//#define PANGO_TYPE_FONT_FACE              (pango_font_face_get_type ())
//#define PANGO_FONT_FACE(object)           (G_TYPE_CHECK_INSTANCE_CAST ((object), PANGO_TYPE_FONT_FACE, PangoFontFace))
//#define PANGO_IS_FONT_FACE(object)        (G_TYPE_CHECK_INSTANCE_TYPE ((object), PANGO_TYPE_FONT_FACE))
//
//GType      pango_font_face_get_type       (void) G_GNUC_CONST;
//
//PangoFontDescription *pango_font_face_describe       (PangoFontFace  *face);
//const char           *pango_font_face_get_face_name  (PangoFontFace  *face) G_GNUC_PURE;
//void                  pango_font_face_list_sizes     (PangoFontFace  *face,
//						      int           **sizes,
//						      int            *n_sizes);
//gboolean              pango_font_face_is_synthesized (PangoFontFace  *face) G_GNUC_PURE;
//
//#ifdef PANGO_ENABLE_BACKEND
//
//#define PANGO_FONT_FACE_CLASS(klass)      (G_TYPE_CHECK_CLASS_CAST ((klass), PANGO_TYPE_FONT_FACE, PangoFontFaceClass))
//#define PANGO_IS_FONT_FACE_CLASS(klass)   (G_TYPE_CHECK_CLASS_TYPE ((klass), PANGO_TYPE_FONT_FACE))
//#define PANGO_FONT_FACE_GET_CLASS(obj)    (G_TYPE_INSTANCE_GET_CLASS ((obj), PANGO_TYPE_FONT_FACE, PangoFontFaceClass))
//
//typedef struct _PangoFontFaceClass   PangoFontFaceClass;
//
///**
// * PangoFontFace:
// *
// * The #PangoFontFace structure is used to represent a group of fonts with
// * the same family, slant, weight, width, but varying sizes.
// */
//struct _PangoFontFace
//{
//  GObject parent_instance;
//};
//
//struct _PangoFontFaceClass
//{
//  GObjectClass parent_class;
//
//  /*< public >*/
//
//  const char           * (*get_face_name)  (PangoFontFace *face);
//  PangoFontDescription * (*describe)       (PangoFontFace *face);
//  void                   (*list_sizes)     (PangoFontFace  *face,
//					    int           **sizes,
//					    int            *n_sizes);
//  gboolean               (*is_synthesized) (PangoFontFace *face);
//
//  /*< private >*/
//
//  /* Padding for future expansion */
//  void (*_pango_reserved3) (void);
//  void (*_pango_reserved4) (void);
//};
//
//#endif /* PANGO_ENABLE_BACKEND */
//
///*
// * PangoFont
// */
//
///**
// * PANGO_TYPE_FONT:
// *
// * The #GObject type for #PangoFont.
// */
///**
// * PANGO_FONT:
// * @object: a #GObject.
// *
// * Casts a #GObject to a #PangoFont.
// */
///**
// * PANGO_IS_FONT:
// * @object: a #GObject.
// *
// * Returns: %TRUE if @object is a #PangoFont.
// */
//#define PANGO_TYPE_FONT              (pango_font_get_type ())
//#define PANGO_FONT(object)           (G_TYPE_CHECK_INSTANCE_CAST ((object), PANGO_TYPE_FONT, PangoFont))
//#define PANGO_IS_FONT(object)        (G_TYPE_CHECK_INSTANCE_TYPE ((object), PANGO_TYPE_FONT))
//
//GType                 pango_font_get_type          (void) G_GNUC_CONST;
//
//PangoFontDescription *pango_font_describe          (PangoFont        *font);
//PangoFontDescription *pango_font_describe_with_absolute_size (PangoFont        *font);
//PangoCoverage *       pango_font_get_coverage      (PangoFont        *font,
//						    PangoLanguage    *language);
//PangoEngineShape *    pango_font_find_shaper       (PangoFont        *font,
//						    PangoLanguage    *language,
//						    guint32           ch);
//PangoFontMetrics *    pango_font_get_metrics       (PangoFont        *font,
//						    PangoLanguage    *language);
//void                  pango_font_get_glyph_extents (PangoFont        *font,
//						    PangoGlyph        glyph,
//						    PangoRectangle   *ink_rect,
//						    PangoRectangle   *logical_rect);
//PangoFontMap         *pango_font_get_font_map      (PangoFont        *font);
//
//#ifdef PANGO_ENABLE_BACKEND
//
//#define PANGO_FONT_CLASS(klass)      (G_TYPE_CHECK_CLASS_CAST ((klass), PANGO_TYPE_FONT, PangoFontClass))
//#define PANGO_IS_FONT_CLASS(klass)   (G_TYPE_CHECK_CLASS_TYPE ((klass), PANGO_TYPE_FONT))
//#define PANGO_FONT_GET_CLASS(obj)    (G_TYPE_INSTANCE_GET_CLASS ((obj), PANGO_TYPE_FONT, PangoFontClass))
//
//typedef struct _PangoFontClass       PangoFontClass;
//
///**
// * PangoFont:
// *
// * The #PangoFont structure is used to represent
// * a font in a rendering-system-independent matter.
// * To create an implementation of a #PangoFont,
// * the rendering-system specific code should allocate
// * a larger structure that contains a nested
// * #PangoFont, fill in the <structfield>klass</structfield> member of
// * the nested #PangoFont with a pointer to
// * a appropriate #PangoFontClass, then call
// * pango_font_init() on the structure.
// *
// * The #PangoFont structure contains one member
// * which the implementation fills in.
// */
//struct _PangoFont
//{
//  GObject parent_instance;
//};
//
//struct _PangoFontClass
//{
//  GObjectClass parent_class;
//
//  /*< public >*/
//
//  PangoFontDescription *(*describe)           (PangoFont      *font);
//  PangoCoverage *       (*get_coverage)       (PangoFont      *font,
//					       PangoLanguage  *lang);
//  PangoEngineShape *    (*find_shaper)        (PangoFont      *font,
//					       PangoLanguage  *lang,
//					       guint32         ch);
//  void                  (*get_glyph_extents)  (PangoFont      *font,
//					       PangoGlyph      glyph,
//					       PangoRectangle *ink_rect,
//					       PangoRectangle *logical_rect);
//  PangoFontMetrics *    (*get_metrics)        (PangoFont      *font,
//					       PangoLanguage  *language);
//  PangoFontMap *        (*get_font_map)       (PangoFont      *font);
//  PangoFontDescription *(*describe_absolute)  (PangoFont      *font);
//  /*< private >*/
//
//  /* Padding for future expansion */
//  void (*_pango_reserved1) (void);
//  void (*_pango_reserved2) (void);
//};
//
///* used for very rare and miserable situtations that we cannot even
// * draw a hexbox
// */
//#define PANGO_UNKNOWN_GLYPH_WIDTH  10
//#define PANGO_UNKNOWN_GLYPH_HEIGHT 14
//
//#endif /* PANGO_ENABLE_BACKEND */
//
///**
// * PANGO_GLYPH_EMPTY:
// *
// * The %PANGO_GLYPH_EMPTY macro represents a #PangoGlyph value that has a
// *  special meaning, which is a zero-width empty glyph.  This is useful for
// * example in shaper modules, to use as the glyph for various zero-width
// * Unicode characters (those passing pango_is_zero_width()).
// */
///**
// * PANGO_GLYPH_INVALID_INPUT:
// *
// * The %PANGO_GLYPH_INVALID_INPUT macro represents a #PangoGlyph value that has a
// * special meaning of invalid input.  #PangoLayout produces one such glyph
// * per invalid input UTF-8 byte and such a glyph is rendered as a crossed
// * box.
// *
// * Note that this value is defined such that it has the %PANGO_GLYPH_UNKNOWN_FLAG
// * on.
// *
// * Since: 1.20
// */
///**
// * PANGO_GLYPH_UNKNOWN_FLAG:
// *
// * The %PANGO_GLYPH_UNKNOWN_FLAG macro is a flag value that can be added to
// * a #gunichar value of a valid Unicode character, to produce a #PangoGlyph
// * value, representing an unknown-character glyph for the respective #gunichar.
// */
///**
// * PANGO_GET_UNKNOWN_GLYPH:
// * @wc: a Unicode character
// *
// * The way this unknown glyphs are rendered is backend specific.  For example,
// * a box with the hexadecimal Unicode code-point of the character written in it
// * is what is done in the most common backends.
// *
// * Returns: a #PangoGlyph value that means no glyph was found for @wc.
// */
//#define PANGO_GLYPH_EMPTY           ((PangoGlyph)0x0FFFFFFF)
//#define PANGO_GLYPH_INVALID_INPUT   ((PangoGlyph)0xFFFFFFFF)
//#define PANGO_GLYPH_UNKNOWN_FLAG    ((PangoGlyph)0x10000000)
//#define PANGO_GET_UNKNOWN_GLYPH(wc) ((PangoGlyph)(wc)|PANGO_GLYPH_UNKNOWN_FLAG)
//
//
