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
	"github.com/conformal/gotk3/glib"
	"unsafe"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.pango_ellipsize_mode_get_type()), marshalEllipsizeMode},
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
