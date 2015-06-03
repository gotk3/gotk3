// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Go bindings for Pango.
package pango

// #cgo pkg-config: pango
// #include <pango/pango.h>
import "C"
import (
	"unsafe"

	"github.com/andre-hub/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.pango_ellipsize_mode_get_type()), marshalEllipsizeMode},
		{glib.Type(C.pango_style_get_type()), marshalStyle},
		{glib.Type(C.pango_weight_get_type()), marshalWeight},
		{glib.Type(C.pango_underline_get_type()), marshalUnderline},
		{glib.Type(C.pango_variant_get_type()), marshalVariant},
		{glib.Type(C.pango_wrap_mode_get_type()), marshalWrapMode},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

// EllipsizeMode is a representation of Pango's PangoEllipsizeMode.
type EllipsizeMode int

const (
	ELLIPSIZE_NONE   EllipsizeMode = C.PANGO_ELLIPSIZE_NONE
	ELLIPSIZE_START  EllipsizeMode = C.PANGO_ELLIPSIZE_START
	ELLIPSIZE_MIDDLE EllipsizeMode = C.PANGO_ELLIPSIZE_MIDDLE
	ELLIPSIZE_END    EllipsizeMode = C.PANGO_ELLIPSIZE_END
)

func marshalEllipsizeMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EllipsizeMode(c), nil
}

// Style is a representation of Pango's PangoStyle
type Style int

const (
	STYLE_NORMAL  Style = C.PANGO_STYLE_NORMAL
	STYLE_OBLIQUE Style = C.PANGO_STYLE_OBLIQUE
	STYLE_ITALIC  Style = C.PANGO_STYLE_ITALIC
)

func marshalStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Style(c), nil
}

// Weight is a representation of Pango's PangoWeight.
type Weight int

const (
	//WEIGHT_THIN       Weight = C.PANGO_WEIGHT_THIN
	WEIGHT_ULTRALIGHT Weight = C.PANGO_WEIGHT_ULTRALIGHT
	WEIGHT_LIGHT      Weight = C.PANGO_WEIGHT_LIGHT
	//WEIGHT_SEMILIGHT  Weight = C.PANGO_WEIGHT_SEMILIGHT
	//WEIGHT_BOOK       Weight = C.PANGO_WEIGHT_BOOK
	WEIGHT_NORMAL Weight = C.PANGO_WEIGHT_NORMAL
	//WEIGHT_MEDIUM     Weight = C.PANGO_WEIGHT_MEDIUM
	WEIGHT_SEMIBOLD  Weight = C.PANGO_WEIGHT_SEMIBOLD
	WEIGHT_BOLD      Weight = C.PANGO_WEIGHT_BOLD
	WEIGHT_ULTRABOLD Weight = C.PANGO_WEIGHT_ULTRABOLD
	WEIGHT_HEAVY     Weight = C.PANGO_WEIGHT_HEAVY
	//WEIGHT_ULTRAHEAVY Weight = C.PANGO_WEIGHT_ULTRAHEAVY
)

func marshalWeight(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Weight(c), nil
}

// Underline is a representation of Pango's PangoUnderline
type Underline int

const (
	UNDERLINE_NONE   Underline = C.PANGO_UNDERLINE_NONE
	UNDERLINE_SINGLE Underline = C.PANGO_UNDERLINE_SINGLE
	UNDERLINE_DOUBLE Underline = C.PANGO_UNDERLINE_DOUBLE
	UNDERLINE_LOW    Underline = C.PANGO_UNDERLINE_LOW
	//UNDERLINE_ERROR  Underline = C.PANGO_UNDERLINE_ERROR
)

func marshalUnderline(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Underline(c), nil
}

// Variant is a representation of Pango's PangoVariant
type Variant int

const (
	VARIANT_NORMAL     Variant = C.PANGO_VARIANT_NORMAL
	VARIANT_SMALL_CAPS Variant = C.PANGO_VARIANT_SMALL_CAPS
)

func marshalVariant(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Variant(c), nil
}

// WrapMode is a representation of Pango's PangoWrapMode.
type WrapMode int

const (
	WRAP_WORD      WrapMode = C.PANGO_WRAP_WORD
	WRAP_CHAR      WrapMode = C.PANGO_WRAP_CHAR
	WRAP_WORD_CHAR WrapMode = C.PANGO_WRAP_WORD_CHAR
)

func marshalWrapMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WrapMode(c), nil
}
