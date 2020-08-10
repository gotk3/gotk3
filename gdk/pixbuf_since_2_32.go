// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4,!gdk_pixbuf_2_6,!gdk_pixbuf_2_8,!gdk_pixbuf_2_12,!gdk_pixbuf_2_14,!gdk_pixbuf_2_22,!gdk_pixbuf_2_24,!gdk_pixbuf_2_26,!gdk_pixbuf_2_28,!gdk_pixbuf_2_30

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// Image Data in Memory

// PixbufNewFromBytes is a wrapper around gdk_pixbuf_new_from_bytes().
// see go package "encoding/base64"
func PixbufNewFromBytes(pixbufData []byte, cs Colorspace, hasAlpha bool, bitsPerSample, width, height, rowStride int) (*Pixbuf, error) {
	arrayPtr := (*C.GBytes)(unsafe.Pointer(&pixbufData[0]))

	c := C.gdk_pixbuf_new_from_bytes(
		arrayPtr,
		C.GdkColorspace(cs),
		gbool(hasAlpha),
		C.int(bitsPerSample),
		C.int(width),
		C.int(height),
		C.int(rowStride),
	)

	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })

	return p, nil
}

// PixbufNewFromBytesOnly is a convenient alternative to PixbufNewFromBytes() and also a wrapper around gdk_pixbuf_new_from_bytes().
// see go package "encoding/base64"
func PixbufNewFromBytesOnly(pixbufData []byte) (*Pixbuf, error) {
	pixbufLoader, err := PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	return pixbufLoader.WriteAndReturnPixbuf(pixbufData)
}

// File loading

// TODO:
// gdk_pixbuf_get_file_info_async().
// gdk_pixbuf_get_file_info_finish().

// The GdkPixbuf Structure

// TODO:
// gdk_pixbuf_get_options().
// gdk_pixbuf_read_pixels().
