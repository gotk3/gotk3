// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0 gmodule-2.0
// #include <glib.h>
// #include <gmodule.h>
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
// #include "pixbuf_since_2_4.go.h"
import "C"
import (
	"errors"
	"io"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/internal/callback"
)

// File saving

//export goPixbufSaveCallback
func goPixbufSaveCallback(buf *C.gchar, count C.gsize, gerr **C.GError, id C.gpointer) C.gboolean {
	v := callback.Get(uintptr(id))

	if v == nil {
		C._pixbuf_error_set_callback_not_found(gerr)
		return C.FALSE
	}

	var bytes []byte
	header := (*reflect.SliceHeader)((unsafe.Pointer(&bytes)))
	header.Cap = int(count)
	header.Len = int(count)
	header.Data = uintptr(unsafe.Pointer(buf))

	_, err := v.(io.Writer).Write(bytes)
	if err != nil {
		cerr := C.CString(err.Error())
		defer C.free(unsafe.Pointer(cerr))

		C._pixbuf_error_set(gerr, cerr)
		return C.FALSE
	}

	return C.TRUE
}

// WritePNG is a convenience wrapper around gdk_pixbuf_save_to_callback() for
// saving images using a streaming callback API. Compression is a number from 0
// to 9.
func (v *Pixbuf) WritePNG(w io.Writer, compression int) error {
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(ccompression))

	id := callback.Assign(w)

	var err *C.GError
	c := C._gdk_pixbuf_save_png_writer(v.native(), C.gpointer(id), &err, ccompression)

	callback.Delete(id)

	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}

	return nil
}

// WriteJPEG is a convenience wrapper around gdk_pixbuf_save_to_callback() for
// saving images using a streaming callback API. Quality is a number from 0 to
// 100.
func (v *Pixbuf) WriteJPEG(w io.Writer, quality int) error {
	cquality := C.CString(strconv.Itoa(quality))
	defer C.free(unsafe.Pointer(cquality))

	id := callback.Assign(w)

	var err *C.GError
	c := C._gdk_pixbuf_save_jpeg_writer(v.native(), C.gpointer(id), &err, cquality)

	callback.Delete(id)

	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}

	return nil
}

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
