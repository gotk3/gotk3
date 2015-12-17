// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
// Copyright (c) 2015 Axel von Blomberg
// This file has been derived from gdk.go
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

// Go bindings for GDK 3 Inmem API
// For documentation see https://developer.gnome.org/gdk-pixbuf/unstable/gdk-pixbuf-Image-Data-in-Memory.html

// +build gdk_3_32
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk_inmem_2_32.h"
import "C"
import (
	"github.com/gotk3/gotk3/glib"
	"runtime"
	"unsafe"
)

// Explicitly hold a reference for each slice that has been converted
// to a GBytes object. The destroy callback function given to GBytes
// constructor removes this reference again.
var glibGBytesReferences map[unsafe.Pointer][]byte

func init_inmem() {
	glibGBytesReferences = make(map[unsafe.Pointer][]byte)
}

//export gdkDestroyGBytes
func gdkDestroyGBytes(data unsafe.Pointer) {
	delete(glibGBytesReferences, data)
}

// PixbufNewFromBytes is a wrapper around gdk_pixbuf_new_from_bytes().
func PixbufNewFromBytes(data []byte, colorspace Colorspace, hasAlpha bool,
	bitsPerSample, width, height, rowstride int) (*Pixbuf, error) {
	glibGBytesReferences[unsafe.Pointer(&data[0])] = data
	gBytes := C.toGBytes(unsafe.Pointer(&data[0]), C.int(len(data)))
	c := C.gdk_pixbuf_new_from_bytes(gBytes, C.GdkColorspace(colorspace),
		gbool(hasAlpha), C.int(bitsPerSample), C.int(width), C.int(height),
		C.int(rowstride))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}
