// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4,!gdk_pixbuf_2_6,!gdk_pixbuf_2_8,!gdk_pixbuf_2_12,!gdk_pixbuf_2_14,!gdk_pixbuf_2_22,!gdk_pixbuf_2_24

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"

// File Loading

// TODO:
// gdk_pixbuf_new_from_resource().
// gdk_pixbuf_new_from_resource_at_scale().

// The GdkPixbuf Structure

// GetByteLength is a wrapper around gdk_pixbuf_get_byte_length().
func (v *Pixbuf) GetByteLength() int {
	c := C.gdk_pixbuf_get_byte_length(v.native())
	return int(c)
}
