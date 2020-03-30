// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4,!gdk_pixbuf_2_6,!gdk_pixbuf_2_8

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

// Utilities

// ApplyEmbeddedOrientation is a wrapper around gdk_pixbuf_apply_embedded_orientation().
func (v *Pixbuf) ApplyEmbeddedOrientation() (*Pixbuf, error) {
	c := C.gdk_pixbuf_apply_embedded_orientation(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}
