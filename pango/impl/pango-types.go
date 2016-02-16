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
// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/pango"
)

// LogAttr is a representation of PangoLogAttr.
type logAttr struct {
	pangoLogAttr *C.PangoLogAttr
}

// Native returns a pointer to the underlying PangoLayout.
func (v *logAttr) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *logAttr) native() *C.PangoLogAttr {
	return (*C.PangoLogAttr)(unsafe.Pointer(v.pangoLogAttr))
}

// EngineLang is a representation of PangoEngineLang.
type engineLang struct {
	pangoEngineLang *C.PangoEngineLang
}

// Native returns a pointer to the underlying PangoLayout.
func (v *engineLang) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *engineLang) native() *C.PangoEngineLang {
	return (*C.PangoEngineLang)(unsafe.Pointer(v.pangoEngineLang))
}

// EngineShape is a representation of PangoEngineShape.
type engineShape struct {
	pangoEngineShape *C.PangoEngineShape
}

// Native returns a pointer to the underlying PangoLayout.
func (v *engineShape) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *engineShape) native() *C.PangoEngineShape {
	return (*C.PangoEngineShape)(unsafe.Pointer(v.pangoEngineShape))
}

// Font is a representation of PangoFont.
type font struct {
	pangoFont *C.PangoFont
}

// Native returns a pointer to the underlying PangoLayout.
func (v *font) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *font) native() *C.PangoFont {
	return (*C.PangoFont)(unsafe.Pointer(v.pangoFont))
}

// FontMap is a representation of PangoFontMap.
type fontMap struct {
	pangoFontMap *C.PangoFontMap
}

// Native returns a pointer to the underlying PangoLayout.
func (v *fontMap) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *fontMap) native() *C.PangoFontMap {
	return (*C.PangoFontMap)(unsafe.Pointer(v.pangoFontMap))
}

// Rectangle is a representation of PangoRectangle.
type rectangle struct {
	pangoRectangle *C.PangoRectangle
}

// Native returns a pointer to the underlying PangoLayout.
func (v *rectangle) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *rectangle) native() *C.PangoRectangle {
	return (*C.PangoRectangle)(unsafe.Pointer(v.pangoRectangle))
}

//void pango_extents_to_pixels (PangoRectangle *inclusive,
//			      PangoRectangle *nearest);
func (inclusive *rectangle) ExtentsToPixels(nearest pango.Rectangle) {
	C.pango_extents_to_pixels(inclusive.native(), toRectangle(nearest).native())
}

func RectangleNew(x, y, width, height int) *rectangle {
	r := new(rectangle)
	r.pangoRectangle = C.createPangoRectangle((C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height))
	return r
}
