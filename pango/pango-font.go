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

type Syle int

const (
	STYLE_NORMAL  Syle = C.PANGO_PANGO_STYLE_NORMAL
	STYLE_OBLIQUE Syle = C.PANGO_PANGO_STYLE_OBLIQUE
	STYLE_ITALIC  Syle = C.PANGO_PANGO_STYLE_ITALIC
)

type Variant int

const (
	VARIANT_NORMAL     Variant = C.PANGO_PANGO_VARIANT_NORMAL
	VARIANT_SMALL_CAPS Variant = C.PANGO_PANGO_VARIANT_SMALL_CAPS
)

type Weight int

const (
	WEIGHT_THIN       Weight = C.PANGO_PANGO_WEIGHT_THIN       /* 100 */
	WEIGHT_ULTRALIGHT Weight = C.PANGO_PANGO_WEIGHT_ULTRALIGHT /* 200 */
	WEIGHT_LIGHT      Weight = C.PANGO_PANGO_WEIGHT_LIGHT      /* 300 */
	WEIGHT_SEMILIGHT  Weight = C.PANGO_PANGO_WEIGHT_SEMILIGHT  /* 350 */
	WEIGHT_BOOK       Weight = C.PANGO_PANGO_WEIGHT_BOOK       /* 380 */
	WEIGHT_NORMAL     Weight = C.PANGO_PANGO_WEIGHT_NORMAL     /* 400 */
	WEIGHT_MEDIUM     Weight = C.PANGO_PANGO_WEIGHT_MEDIUM     /* 500 */
	WEIGHT_SEMIBOLD   Weight = C.PANGO_PANGO_WEIGHT_SEMIBOLD   /* 600 */
	WEIGHT_BOLD       Weight = C.PANGO_PANGO_WEIGHT_BOLD       /* 700 */
	WEIGHT_ULTRABOLD  Weight = C.PANGO_PANGO_WEIGHT_ULTRABOLD  /* 800 */
	WEIGHT_HEAVY      Weight = C.PANGO_PANGO_WEIGHT_HEAVY      /* 900 */
	WEIGHT_ULTRAHEAVY Weight = C.PANGO_PANGO_WEIGHT_ULTRAHEAVY /* 1000 */

)

type Stretch int

const (
	STRETCH_ULTRA_CONDENSED        Stretch = C.PANGO_PANGO_STRETCH_ULTRA_CONDENSED
	STRETCH_EXTRA_CONDENSEDStretch         = C.PANGO_PANGO_STRETCH_EXTRA_CONDENSED
	STRETCH_CONDENSEDStretch               = C.PANGO_PANGO_STRETCH_CONDENSED
	STRETCH_SEMI_CONDENSEDStretch          = C.PANGO_PANGO_STRETCH_SEMI_CONDENSED
	STRETCH_NORMALStretch                  = C.PANGO_PANGO_STRETCH_NORMAL
	STRETCH_SEMI_EXPANDEDStretch           = C.PANGO_PANGO_STRETCH_SEMI_EXPANDED
	STRETCH_EXPANDEDStretch                = C.PANGO_PANGO_STRETCH_EXPANDED
	STRETCH_EXTRA_EXPANDEDStretch          = C.PANGO_PANGO_STRETCH_EXTRA_EXPANDED
	STRETCH_ULTRA_EXPANDEDStretch          = C.PANGO_PANGO_STRETCH_ULTRA_EXPANDED
)

type FontMask int

const (
	FONT_MASK_FAMILY          FontMask = C.PANGO_PANGO_FONT_MASK_FAMILY  /*  1 << 0 */
	FONT_MASK_STYLEFontMask            = C.PANGO_PANGO_FONT_MASK_STYLE   /*  1 << 1 */
	FONT_MASK_VARIANTFontMask          = C.PANGO_PANGO_FONT_MASK_VARIANT /*  1 << 2 */
	FONT_MASK_WEIGHTFontMask           = C.PANGO_PANGO_FONT_MASK_WEIGHT  /*  1 << 3 */
	FONT_MASK_STRETCHFontMask          = C.PANGO_PANGO_FONT_MASK_STRETCH /*  1 << 4 */
	FONT_MASK_SIZEFontMask             = C.PANGO_PANGO_FONT_MASK_SIZE    /*  1 << 5 */
	FONT_MASK_GRAVITYFontMask          = C.PANGO_PANGO_FONT_MASK_GRAVITY /*  1 << 6 */
)
