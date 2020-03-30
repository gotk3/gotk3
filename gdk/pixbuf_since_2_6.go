// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4

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

// File Loading

// PixbufNewFromFileAtScale is a wrapper around gdk_pixbuf_new_from_file_at_scale().
func PixbufNewFromFileAtScale(filename string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError = nil
	c := C.gdk_pixbuf_new_from_file_at_scale(cstr, C.int(width), C.int(height),
		gbool(preserveAspectRatio), &err)
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

// Scaling

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (v *Pixbuf) RotateSimple(angle PixbufRotation) (*Pixbuf, error) {
	c := C.gdk_pixbuf_rotate_simple(v.native(), C.GdkPixbufRotation(angle))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (v *Pixbuf) Flip(horizontal bool) (*Pixbuf, error) {
	c := C.gdk_pixbuf_flip(v.native(), gbool(horizontal))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}
