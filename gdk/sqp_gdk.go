package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
/*

*/
import "C"
import (
	"github.com/andre-hub/gotk3/glib"
	"unsafe"
)

func (v *EventKey) State() uint {
	c := v.native().state
	return uint(c)
}

type RGBA struct {
	RGBA C.GdkRGBA
}

func NewRGBA(values ...float64) *RGBA {
	c := &RGBA{}
	if len(values) > 0 {
		c.RGBA.red = C.gdouble(values[0])
	}
	if len(values) > 1 {
		c.RGBA.green = C.gdouble(values[1])
	}
	if len(values) > 2 {
		c.RGBA.blue = C.gdouble(values[2])
	}
	if len(values) > 3 {
		c.RGBA.alpha = C.gdouble(values[3])
	}
	return c
}

func (c *RGBA) Floats() []float64 {
	return []float64{float64(c.RGBA.red), float64(c.RGBA.green), float64(c.RGBA.blue), float64(c.RGBA.alpha)}
}

func (v *RGBA) Native() uintptr {
	return uintptr(unsafe.Pointer(&v.RGBA))
}

// PixbufGetType is a wrapper around gdk_pixbuf_get_type().
func PixbufGetType() glib.Type {
	return glib.Type(C.gdk_pixbuf_get_type())
}
