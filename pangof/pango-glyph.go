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

package pangof

// #cgo pkg-config: pango
// #include <pango/pango.h>
// #include "pango.go.h"
// #include <stdlib.h>
import "C"
import
//	"github.com/andre-hub/gotk3/glib"
//	"github.com/andre-hub/gotk3/cairo"
"unsafe"

// GlyphGeometry is a representation of PangoGlyphGeometry.
type glyphGeometry struct {
	pangoGlyphGeometry *C.PangoGlyphGeometry
}

// Native returns a pointer to the underlying PangoLayout.
func (v *glyphGeometry) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *glyphGeometry) native() *C.PangoGlyphGeometry {
	return (*C.PangoGlyphGeometry)(unsafe.Pointer(v.pangoGlyphGeometry))
}

// GlyphVisAttr is a representation of PangoGlyphVisAttr.
type glyphVisAttr struct {
	pangoGlyphVisAttr *C.PangoGlyphGeometry
}

// Native returns a pointer to the underlying PangoGlyphVisAttr.
func (v *glyphVisAttr) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *glyphVisAttr) native() *C.PangoGlyphVisAttr {
	return (*C.PangoGlyphVisAttr)(unsafe.Pointer(v.pangoGlyphVisAttr))
}

// GlyphInfo is a representation of PangoGlyphInfo.
type glyphInfo struct {
	pangoGlyphInfo *C.PangoGlyphInfo
}

// Native returns a pointer to the underlying PangoGlyphInfo.
func (v *glyphInfo) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *glyphInfo) native() *C.PangoGlyphInfo {
	return (*C.PangoGlyphInfo)(unsafe.Pointer(v.pangoGlyphInfo))
}

// GlyphGeometry is a representation of PangoGlyphString.
type glyphString struct {
	pangoGlyphString *C.PangoGlyphString
}

// Native returns a pointer to the underlying PangoGlyphString.
func (v *glyphString) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *glyphString) native() *C.PangoGlyphString {
	return (*C.PangoGlyphString)(unsafe.Pointer(v.pangoGlyphString))
}
