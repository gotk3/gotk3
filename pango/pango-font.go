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
