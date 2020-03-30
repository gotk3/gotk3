// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// File saving

// TODO:
// GdkPixbufSaveFunc
// gdk_pixbuf_save_to_callback().
// gdk_pixbuf_save_to_callbackv().
// gdk_pixbuf_save_to_buffer().
// gdk_pixbuf_save_to_bufferv().

// File Loading

// PixbufNewFromFileAtSize is a wrapper around gdk_pixbuf_new_from_file_at_size().
func PixbufNewFromFileAtSize(filename string, width, height int) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError = nil
	c := C.gdk_pixbuf_new_from_file_at_size(cstr, C.int(width), C.int(height), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// PixbufGetFileInfo is a wrapper around gdk_pixbuf_get_file_info().
func PixbufGetFileInfo(filename string) (*PixbufFormat, int, int, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var cw, ch C.gint
	format := C.gdk_pixbuf_get_file_info((*C.gchar)(cstr), &cw, &ch)
	if format == nil {
		return nil, -1, -1, nilPtrErr
	}
	// The returned PixbufFormat value is owned by Pixbuf and should not be freed.
	return wrapPixbufFormat(format), int(cw), int(ch), nil
}
